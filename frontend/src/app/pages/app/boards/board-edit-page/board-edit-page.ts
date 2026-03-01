import { AsyncPipe } from '@angular/common';
import { Component, inject } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { BoardModel } from '@entities/board';
import { BoardDeleteFeature, BoardEditFeature } from '@root/src/app/features';
import { UiService } from '@shared/ui';
import { filter, map, Observable, tap } from 'rxjs';

@Component({
    selector: 'app-board-edit-page',
    imports: [BoardEditFeature, BoardDeleteFeature, AsyncPipe],
    template: `
        @if (board$ | async; as board) {
            <h4>Board "{{ board.title }}" Edit</h4>
            <app-board-edit-feature [board]="board" />
            <hr />
            <app-board-delete-feature [boardId]="board.id" [boardTitle]="board.title" />
        }
    `,
})
export class BoardEditPage {
    private ui = inject(UiService);

    route = inject(ActivatedRoute);
    projectId: number;
    board$: Observable<BoardModel> = this.route.parent!.data.pipe(
        filter((data) => !!data['board']),
        map((data) => data['board'] as BoardModel),
        tap((board) => this.ui.setPageTitle(`Edit board: ${board.title}`)),
    );
}
