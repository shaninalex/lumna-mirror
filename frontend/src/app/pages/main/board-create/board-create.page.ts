import { Component, inject } from '@angular/core';
import { MainLayout } from '@core/layout';
import { AppRoutes } from '@core/routes';
import { BoardCreateFeature } from "@features";

@Component({
    selector: 'lu-board-create-page',
    imports: [MainLayout, BoardCreateFeature],
    templateUrl: './board-create.page.html',
})
export class BoardCreatePage {
    private appRoutes = inject(AppRoutes);
}
