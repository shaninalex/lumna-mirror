import {Component, inject} from '@angular/core';
import {UiService} from '@client/shared/ui';

@Component({
    selector: "jr-page-home",
    template: `Home page`
})
export class PageHomeComponent {
    private ui: UiService = inject(UiService)

    constructor() {
        this.ui.setTitle("Home");
    }
}
