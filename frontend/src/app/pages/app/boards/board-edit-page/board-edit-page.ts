import { AsyncPipe } from '@angular/common';
import { Component, inject } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { BoardModel, BoardState } from '@entities/board';
import { selectBoardById } from '@entities/board/model/board.selectors';
import { Store } from '@ngrx/store';
import { BoardDeleteFeature, BoardEditFeature } from '@root/src/app/features';
import { UiService } from '@shared/ui';
import { filter, Observable, switchMap, tap } from 'rxjs';

@Component({
    selector: 'app-board-edit-page',
    imports: [BoardEditFeature, BoardDeleteFeature, AsyncPipe],
    template: `
        @if (board$ | async; as board) {
            <h1>Board "{{ board.title }}" Edit</h1>
            <app-board-edit-feature [board]="board" />
            <div class="my-4 border-b border-gray-200"></div>
            <app-board-delete-feature [boardTitle]="board.title" [projectId]="projectId" [boardId]="board.id" />
        }
    `,
})
export class BoardEditPage {
    private store = inject(Store<BoardState>);
    private ui = inject(UiService);

    route = inject(ActivatedRoute);
    projectId: string;
    board$: Observable<BoardModel> = this.route.params.pipe(
        tap((params) => (this.projectId = params['id'])),
        switchMap((params) =>
            this.store.select(selectBoardById(params['boardId'])).pipe(
                filter((board) => !!board),
                tap((board) => {
                    const title = `Edit board: ${board.title}`;
                    this.ui.setPageTitle(title);
                }),
            ),
        ),
    );
}
