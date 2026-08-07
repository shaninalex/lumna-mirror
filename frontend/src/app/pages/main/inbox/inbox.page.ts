import { Component, inject } from '@angular/core';
import { MainLayout } from '@core/layout';
import { UiService } from '@shared/ui';

@Component({
    selector: 'lu-inbox-page',
    imports: [MainLayout],
    template: `
        <lu-main-layout>
            <p>inbox works!</p>
        </lu-main-layout>
    `,
})
export class InboxPage {
    private ui = inject(UiService);

    ngOnInit(): void {
        this.ui.setPageTitle("Inbox");
    }
}
