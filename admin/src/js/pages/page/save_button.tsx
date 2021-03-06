import msx from 'lib/msx';
import * as API from 'lib/api';
import Page from 'lib/page';
import * as Toaster from 'components/toaster';

interface PageSaveButtonAttrs {
  page: Page;
  onsave: (page: Page) => void;
  classes?: string;
}

export default class PageSaveButtonComponent {
  private readonly _attrs: PageSaveButtonAttrs;

  constructor(private onsave: (page: Page) => void) {
  }

  save(page: Page) {
    page.save().then((p: API.Page) => {
      page.uuid = p.uuid;
      window.history.replaceState(null, page.title, `/admin/pages/${p.uuid}`);
      return page.saveRoutes();
    })
      .then(() => {
        Toaster.add('Page successfully saved');
        this.onsave(page);
      })
      .catch((err: any) => {
        if (err.detail) {
          Toaster.add(err.detail, 'error');
        } else {
          Toaster.add('Internal server error.', 'error');
        }
      });
  }

  static oninit(v: Mithril.Vnode<PageSaveButtonAttrs, PageSaveButtonComponent>) {
    v.state = new PageSaveButtonComponent(v.attrs.onsave);
  }

  static view({ attrs: { page, classes }, state }: Mithril.Vnode<PageSaveButtonAttrs, PageSaveButtonComponent>) {
    return <a
      class={`button button--green ${classes || ''}`}
      onclick={(e: Event) => { e.stopPropagation(); state.save(page); }}
    >
      Save
    </a>;
  }
}
