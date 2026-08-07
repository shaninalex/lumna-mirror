import { Component, inject } from '@angular/core';
import { RouterLink } from "@angular/router";
import { GlobalLayout } from '@core/layout';
import { UiService } from '@shared/ui';

@Component({
    selector: 'lu-workspaces-page',
    imports: [RouterLink, GlobalLayout],
    templateUrl: 'workspaces.page.html',
})
export class WorkspacesPage {
    private ui = inject(UiService);

    ngOnInit(): void {
        this.ui.setPageTitle("Workspaces")
    }
}
