import { Component, inject } from '@angular/core';
import { MainLayout } from '@core/layout';
import { UiService } from '@shared/ui';
import { RouterLink } from "@angular/router";

@Component({
    selector: 'lu-backlog',
    imports: [MainLayout, RouterLink],
    templateUrl: './backlog.page.html',
})
export class BacklogPage {
    private ui = inject(UiService);

    ngOnInit(): void {
        this.ui.setPageTitle("Backlog");
    }
}
