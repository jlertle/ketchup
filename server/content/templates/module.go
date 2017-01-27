package templates

import (
	"io"
	"mime"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/octavore/naga/service"
	"github.com/octavore/nagax/logger"

	"github.com/octavore/press/proto/press/models"
	"github.com/octavore/press/server/config"
	"github.com/octavore/press/util/errors"
)

const themeDir = "themes"

type ThemesConfig struct {
	Themes struct {
		Path        string `json:"dir"`
		RegistryURL string `json:"registry_url"`
	} `json:"themes"`
}

// ThemeStore is a interface to support multiple theme backends.
type ThemeStore interface {
	List() ([]*models.Theme, error)
	Add(io.Reader) (*models.Theme, error)
	Get(string) (*models.Theme, error)
	GetTemplate(t *models.Theme, template string) (*models.ThemeTemplate, error)
	GetAsset(t *models.Theme, asset string) (*models.ThemeAsset, error)
}

// Module templates provides support for looking up and using themes and
// their corresponding templates.
type Module struct {
	ConfigModule *config.Module
	Logger       *logger.Module

	Stores []ThemeStore
	config ThemesConfig
}

// Init implements service.Init
func (m *Module) Init(c *service.Config) {
	c.Setup = func() error {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}

		err = m.ConfigModule.ReadConfig(&m.config)
		if err != nil {
			return err
		}

		m.config.Themes.Path = m.ConfigModule.DataPath(m.config.Themes.Path, themeDir)
		dataDir := path.Join(wd, m.config.Themes.Path)
		m.Stores = []ThemeStore{&defaultStore{}, m.newFileStore(dataDir)}
		return nil
	}
}

// ServeHTTP serves theme assets
func (m *Module) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	asset, err := m.GetAsset(req.URL.Path)
	if err != nil {
		m.Logger.Warning(err)
		http.NotFound(rw, req)
		return
	}
	if asset == nil {
		http.NotFound(rw, req)
		return
	}
	ext := path.Ext(req.URL.Path)
	rw.Header().Add("Content-Type", mime.TypeByExtension(ext))
	rw.Write([]byte(asset.GetData()))
}

func (m *Module) getTheme(name string) (ThemeStore, *models.Theme, error) {
	for i := len(m.Stores) - 1; i > -1; i-- {
		store := m.Stores[i]
		theme, err := store.Get(name)
		if err != nil {
			return nil, nil, errors.Wrap(err)
		}
		if theme != nil {
			theme.Name = &name
			return store, theme, nil
		}
	}
	return nil, nil, nil
}

// GetTheme returns a theme with all related assets populated.
func (m *Module) GetTheme(name string) (*models.Theme, error) {
	_, theme, err := m.getTheme(name)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	return theme, nil
}

func (m *Module) GetTemplate(theme, template string) (*models.ThemeTemplate, error) {
	tmpl, err := m.getTemplate(theme, template)
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}

// ListThemes returns a list of all known themes.
func (m *Module) ListThemes() ([]*models.Theme, error) {
	themes := []*models.Theme{}
	for _, store := range m.Stores {
		t, err := store.List()
		if err != nil {
			return nil, err
		}
		themes = append(themes, t...)
	}
	return themes, nil
}

// GetAsset searches all themes for the named asset
func (m *Module) GetAsset(name string) (*models.ThemeAsset, error) {
	for i := len(m.Stores) - 1; i != 0; i-- {
		store := m.Stores[i]
		themes, err := store.List()
		if err != nil {
			return nil, err
		}
		for _, theme := range themes {
			asset, err := store.GetAsset(theme, name)
			if err != nil {
				return nil, err
			}
			if asset != nil {
				return asset, nil
			}
		}
	}
	return nil, nil
}

func (m *Module) newFileStore(dataDir string) *FileStore {
	f := &FileStore{dataDir: dataDir}
	go func() {
		for range time.Tick(10 * time.Second) {
			err := f.updateThemeDirMap()
			if err != nil {
				m.Logger.Error(err)
			}
		}
	}()
	return f
}
