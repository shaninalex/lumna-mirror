import { AsyncPipe } from '@angular/common';
import { Component, inject } from '@angular/core';
import { BoardListItemComponent } from '@entities/board';
import { selectBoardsByProjectId } from '@entities/board/model/board.selectors';
import { selectCurrentProjectId } from '@entities/project';
import { Store } from '@ngrx/store';
import { filter, switchMap } from 'rxjs';

@Component({
    selector: 'lu-board-list-feature',
    imports: [AsyncPipe, BoardListItemComponent],
    template: `
        <div class="list-group">
            @if (boards$ | async; as boards) {
                @for (board of boards; track $index) {
                    <lu-board-list-item [board]="board" />
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
