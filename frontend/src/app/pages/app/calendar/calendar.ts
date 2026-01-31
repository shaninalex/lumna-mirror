import { Component, inject } from '@angular/core';
import { UiService } from '@shared/ui';

@Component({
    selector: 'app-calendar',
    template: ` <div class="bg-gray-100 p-10 rounded-2xl">CALENDAR NOT IMPLEMENTED YET</div> `,
})
export class Calendar {
    private ui = inject(UiService);

    constructor() {
        this.ui.setPageTitle(`Calendar`);
    }
}
