import { Component, inject, OnInit } from '@angular/core';
import { actionBoardGet, BoardModel, BoardState } from '@entities/board';
import { ActivatedRoute, RouterLink } from '@angular/router';
import { CdkMenu, CdkMenuTrigger } from '@angular/cdk/menu';
import { KanbanBoardFeature } from '@features/kanban-board';
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
                    <button [cdkMenuTriggerFor]="menu" class="cursor-pointer">
                        <i class="fa-solid fa-ellipsis"></i>
                    </button>

                    <ng-template #menu>
                        <div class="dropdown p-4" cdkMenu>
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
    board$: Observable<BoardModel | null>;

    ngOnInit() {
        this.board$ = this.route.params.pipe(
            switchMap((params) =>
                this.store.select(selectBoardById(params['id'])).pipe(
                    // filter((board) => !!board),
                    tap((board) => {
                        if (board) {
                            this.ui.setPageTitle(`Board: ${board.title}`);
                        } else {
                            this.store.dispatch(actionBoardGet({ boardId: params['id'] }));
                        }
                    }),
                ),
            ),
        );
    }
}
