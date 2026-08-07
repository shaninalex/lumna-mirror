import { Component, inject } from '@angular/core';
import { MainLayout } from '@core/layout';
import { UiService } from '@shared/ui';
import { RouterLink } from "@angular/router";

@Component({
    selector: 'lu-board-page',
    imports: [MainLayout, RouterLink],
    templateUrl: './board.page.html',
})
export class BoardPage {
    private ui = inject(UiService);

    ngOnInit(): void {
        this.ui.setPageTitle("Development board")
    }
}
