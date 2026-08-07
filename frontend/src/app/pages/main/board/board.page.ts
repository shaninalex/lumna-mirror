import { Component, inject } from '@angular/core';
import { MainLayout } from '@core/layout';
import { UiService } from '@shared/ui';

@Component({
    selector: 'lu-board-page',
    imports: [MainLayout],
    templateUrl: './board.page.html',
})
export class BoardPage {
    private ui = inject(UiService);

    ngOnInit(): void {
        this.ui.setPageTitle("Development board")
    }
}
