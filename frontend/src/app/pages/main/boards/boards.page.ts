import type { OnInit } from '@angular/core';
import { Component, inject } from '@angular/core';
import { MainLayout } from '@core/layout';
import { UiService } from '@shared/ui';
import { RouterLink } from "@angular/router";
import { AppRoutes } from '@core/routes';
import { Store } from '@ngrx/store';
import { selectCurrentProjectId } from '@entities/project';
import { filter, switchMap } from 'rxjs';
import { selectBoardsByProjectId } from '@entities/board/model/board.selectors';
import { AsyncPipe } from '@angular/common';

@Component({
    selector: 'lu-boards-page',
    imports: [MainLayout, RouterLink, AsyncPipe],
    templateUrl: './boards.page.html',
})
export class BoardsPage implements OnInit {
    private ui = inject(UiService);
    private store = inject(Store);
    
    readonly appRoutes = inject(AppRoutes);

    boards$ = this.store.select(selectCurrentProjectId).pipe(
        filter((projectId) => projectId !== null),
        switchMap((projectId) => {
            console.log('projectId: ', projectId);
            return this.store.select(selectBoardsByProjectId(projectId))
        })
    )

    ngOnInit(): void {
        this.ui.setPageTitle("Boards")
    }
}
