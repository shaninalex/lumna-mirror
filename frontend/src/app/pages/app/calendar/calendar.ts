import { Component, inject } from '@angular/core';
import { UiService } from '@shared/ui';

@Component({
    selector: 'app-calendar',
    template: `
        <div class="alert alert-info" role="alert">
            <h4 class="alert-heading">Important!</h4>
            <hr>
            <p>The page content is not implemented yet.</p>
        </div>
    `,
})
export class Calendar {
    private ui = inject(UiService);

    constructor() {
        this.ui.setPageTitle(`Calendar`);
    }
}
