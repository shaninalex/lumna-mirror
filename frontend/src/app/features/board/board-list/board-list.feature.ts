import { AsyncPipe } from '@angular/common';
import { Component, inject } from '@angular/core';
import { selectBoardsByProjectId } from '@entities/board/model/board.selectors';
import { selectCurrentProjectId } from '@entities/project';
import { Store } from '@ngrx/store';
import { filter, switchMap } from 'rxjs';

@Component({
    selector: 'lu-board-list-feature',
    imports: [AsyncPipe],
    template: `
        <div class="list-group">
            @if (boards$ | async; as boards) {
                @for (board of boards; track $index) {
                    <a
                        routerLink="/app/w/1/board"
                        class="list-group-item list-group-item-action py-3"
                    >
                        <div class="d-flex justify-content-between align-items-start">
                            <div class="me-3">
                                <h5 class="mb-1">{{ board.title }}</h5>

                                <p class="mb-2 text-muted">
                                    Main board for feature development, bugs and improvements.
                                </p>

                                <div class="d-flex gap-3 small text-muted">
                                    <span>42 cards</span>
                                    <span>5 columns</span>
                                    <span>Updated 2 hours ago</span>
                                </div>
                            </div>

                            <i class="fa-solid fa-chevron-right text-muted mt-1"></i>
                        </div>
                    </a>
                }
            }
        </div>
    `,
})
export class BoardListFeature {
    private store = inject(Store);
    boards$ = this.store.select(selectCurrentProjectId).pipe(
        filter((projectId) => projectId !== null),
        switchMap((projectId) => this.store.select(selectBoardsByProjectId(projectId))),
    );
}
