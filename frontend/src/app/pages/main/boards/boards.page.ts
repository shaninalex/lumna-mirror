import { Component, inject } from '@angular/core';
import { MainLayout } from '@core/layout';
import { UiService } from '@shared/ui';
import { RouterLink } from "@angular/router";

@Component({
    selector: 'lu-boards-page',
    imports: [MainLayout, RouterLink],
    templateUrl: './boards.page.html',
})
export class BoardsPage {
    private ui = inject(UiService);

    ngOnInit(): void {
        this.ui.setPageTitle("Boards")
    }
}
