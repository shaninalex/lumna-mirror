import { Component, inject } from '@angular/core';
import { BoardModel } from '@entities/board';
import { ActivatedRoute, RouterLink, RouterOutlet } from '@angular/router';
import {CdkMenu , CdkMenuTrigger} from '@angular/cdk/menu';
import { KanbanBoardFeature } from '@features/kanban-board';
import { filter, map, Observable, tap } from 'rxjs';
import { AsyncPipe } from '@angular/common';
import { UiService } from '@shared/ui';

@Component({
    selector: 'app-board-page',
    imports: [CdkMenu, CdkMenuTrigger, RouterLink, RouterOutlet, KanbanBoardFeature, AsyncPipe],
    template: `
        @if (board$ | async; as board) {
            <div class="card mb-4 bg-body-tertiary">
                <div class="card-body d-flex justify-content-between align-items-center">
                    <div class="fw-bold">{{ board.title }}</div>
                    <div>
                        <button [cdkMenuTriggerFor]="menu" class="btn btn-sm">
                            <i class="fa-solid fa-ellipsis"></i>
                        </button>

                        <ng-template #menu>
                            <div class="dropdown-menu d-block" cdkMenu>
                                <a [routerLink]="['/projects', board.project_id, 'boards', board.id, 'edit']">Edit</a>
                            </div>
                        </ng-template>
                    </div>
                </div>
            </div>

            <app-kanban-board-feature [board]="board"/>
        }
        <router-outlet/>
    `,
})
export class BoardPage {
    private route = inject(ActivatedRoute);
    private ui = inject(UiService);

    board$: Observable<BoardModel> = this.route.data.pipe(
        filter((data) => !!data['board']),
        map((data) => data['board'] as BoardModel),
        tap((board) => this.ui.setPageTitle(`Board: ${board.title}`)),
    );
}
