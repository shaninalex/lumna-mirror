import { Component, inject } from '@angular/core';
import { MainLayout } from '@core';
import { UiService } from '@shared/ui';

@Component({
    selector: 'app-calendar',
    imports: [MainLayout],
    template: `
        <app-main-layout pageTitle="Calendar">
            <div class="bg-gray-100 p-10 rounded-2xl">CALENDAR NOT IMPLEMENTED YET</div>
        </app-main-layout>
    `,
})
export class Calendar {
    private ui = inject(UiService);

    constructor() {
        this.ui.setPageTitle(`Calendar`);
    }
}
