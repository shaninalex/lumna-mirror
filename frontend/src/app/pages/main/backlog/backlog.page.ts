import { Component, inject } from '@angular/core';
import { MainLayout } from '@core/layout';
import { UiService } from '@shared/ui';

@Component({
    selector: 'lu-backlog',
    imports: [MainLayout],
    templateUrl: './backlog.page.html',
})
export class BacklogPage {
    private ui = inject(UiService);

    ngOnInit(): void {
        this.ui.setPageTitle("Backlog");
    }
}
