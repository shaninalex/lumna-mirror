import { Component, inject, OnInit, signal } from '@angular/core';
import { BoardModel, BoardState } from '@entities/board';
import { ActivatedRoute, RouterLink } from '@angular/router';
import { CdkMenu, CdkMenuTrigger } from '@angular/cdk/menu';
import { KanbanBoardFeature } from '@features/kanban-board/kanban-board';
import { UiService } from '@shared/ui';
import { filter, Observable, switchMap, tap } from 'rxjs';
import { AsyncPipe } from '@angular/common';
import { Store } from '@ngrx/store';
import { selectBoardById } from '@entities/board/model/board.selectors';

@Component({
    selector: 'app-board-page',
    imports: [CdkMenu, CdkMenuTrigger, RouterLink, KanbanBoardFeature, AsyncPipe],
    template: `
        @if (board$ | async; as board) {
            <div class="card bg-lime-200 dark:bg-lime-800 flex items-top justify-between mb-4">
                <div>
                    <div class="font-medium">{{ board.title }}</div>
                </div>

                <div>
                    <button [cdkMenuTriggerFor]="menu">
                        <i class="fa-solid fa-ellipsis"></i>
                    </button>

                    <ng-template #menu>
                        <div class="bg-white border border-gray-200 rounded-xl p-4" cdkMenu>
                            <a [routerLink]="['edit']">Edit</a>
                        </div>
                    </ng-template>
                </div>
            </div>

            <app-kanban-board-feature [board]="board" />
        }
    `,
})
export class BoardPage implements OnInit {
    private route = inject(ActivatedRoute);
    private ui = inject(UiService);
    private store = inject(Store<BoardState>);
    board$: Observable<BoardModel>;

    ngOnInit() {
        this.board$ = this.route.params.pipe(
            switchMap((params) =>
                this.store.select(selectBoardById(params['boardId'])).pipe(
                    filter((board) => !!board),
                    tap((board) => this.ui.setPageTitle(`Board: ${board.title}`)),
                ),
            ),
        );
    }
}
