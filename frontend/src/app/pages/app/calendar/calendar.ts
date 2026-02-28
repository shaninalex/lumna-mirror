import { Component, inject } from '@angular/core';
import { UiService } from '@shared/ui';

@Component({
    selector: 'app-calendar',
    template: `
        <article class="message is-info">
            <div class="message-body">
                NOT IMPLEMENTED
            </div>
        </article>
    `,
})
export class Calendar {
    private ui = inject(UiService);

    constructor() {
        this.ui.setPageTitle(`Calendar`);
    }
}
