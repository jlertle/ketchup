package tls

import (
	"io/ioutil"
	"os"
	"path"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xenolf/lego/acme"
)

func TestObtainCert(t *testing.T) {
	module, _ := setup(t)

	err := module.ObtainCert("test@example.com", "example.com")
	assert.NoError(t, err)

	r, err := module.GetRegistration("example.com", false)
	assert.NoError(t, err)
	assert.Equal(t, &Registration{
		Email:    "test@example.com",
		Domain:   "example.com",
		AgreedOn: "2017-01-01T00:51:26Z",
	}, r)

	cr, err := module.LoadCertResource("example.com")
	assert.NoError(t, err)
	assert.Equal(t, &acme.CertificateResource{Domain: "example.com"}, cr)
}

func TestRenew(t *testing.T) {
	module, _ := setup(t)
	err := module.ObtainCert("test@example.com", "example.com")
	assert.NoError(t, err)

	registrationPath, _ := module.getCurrentRegistrationPath("example.com")
	assert.Equal(t, "example.com-2017-01-01-v000.json", path.Base(registrationPath))

	err = module.renewExpiredCerts()
	registrationPath, _ = module.getCurrentRegistrationPath("example.com")
	assert.NoError(t, err)
	assert.Equal(t, "example.com-2017-01-01-v000.json", path.Base(registrationPath))

	module.renewWithinInterval = time.Hour * 24 * 180
	err = module.renewExpiredCerts()
	registrationPath, _ = module.getCurrentRegistrationPath("example.com")
	assert.NoError(t, err)
	assert.Equal(t, "example.com-2017-01-01-v001.json", path.Base(registrationPath))

}

func TestSaveRegistration(t *testing.T) {
	module, dir := setup(t)
	_ = ioutil.WriteFile(
		path.Join(dir, tlsDir, "example.com-2017-01-01-v000.json"), []byte(`{
			"email": "oldadmin@example.com",
			"domain": "example.com"
		}`),
		os.ModePerm,
	)
	err := module.SaveRegistration(&Registration{
		Email:  "admin@example.com",
		Domain: "example.com",
	})
	assert.NoError(t, err)

	expectedPath := path.Join(dir, tlsDir, "example.com-2017-01-01-v001.json")
	_, err = os.Stat(expectedPath)
	assert.NoError(t, err)

	r, err := module.GetRegistration("example.com", false)
	assert.NoError(t, err)
	assert.Equal(t, &Registration{
		Email:  "admin@example.com",
		Domain: "example.com",
	}, r)
}

func TestGetRegistration(t *testing.T) {
	module, dir := setup(t)
	_ = ioutil.WriteFile(
		path.Join(dir, tlsDir, "example.com-2017-01-01-v000.json"), []byte(`{
			"email": "oldadmin@example.com"
		}`),
		os.ModePerm,
	)
	_ = ioutil.WriteFile(
		path.Join(dir, tlsDir, "example.com-2017-01-01-v001.json"), []byte(`{
			"email": "admin@example.com"
		}`),
		os.ModePerm,
	)

	r, err := module.GetRegistration("fakedomain.com", false)
	if assert.NoError(t, err) {
		assert.Nil(t, r)
	}

	r, err = module.GetRegistration("example.com", false)
	if assert.NoError(t, err) {
		assert.Equal(t, &Registration{Email: "admin@example.com"}, r)
	}
}
