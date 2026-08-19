import type { OnInit } from '@angular/core';
import { Component, inject } from '@angular/core';
import { MainLayout } from '@core/layout';
import { UiService } from '@shared/ui';
import { RouterLink } from "@angular/router";
import { AppRoutes } from '@core/routes';

@Component({
    selector: 'lu-boards-page',
    imports: [MainLayout, RouterLink],
    templateUrl: './boards.page.html',
})
export class BoardsPage implements OnInit {
    private ui = inject(UiService);
    readonly appRoutes = inject(AppRoutes);

    ngOnInit(): void {
        this.ui.setPageTitle("Boards")
    }
}
