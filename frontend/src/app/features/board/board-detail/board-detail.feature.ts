import { AsyncPipe } from '@angular/common';
import type { OnInit } from '@angular/core';
import { Component, inject, input } from '@angular/core';
import { Store } from '@ngrx/store';
import type { ColumnModel } from '@entities/column';
import { actionsColumns, ColumnItemComponent, NewColumnFormComponent, selectStatuses } from '@entities/column';
import { filter, type Observable } from 'rxjs';
import { TimeAgoPipe } from '@shared/utils';
import { selectBoard, type BoardModel } from '@entities/board';

@Component({
    selector: 'lu-board-detail-feature',
    imports: [AsyncPipe, NewColumnFormComponent, TimeAgoPipe, ColumnItemComponent],
    templateUrl: './board-detail.feature.html',
})
export class BoardDetailFeature implements OnInit {
    private store = inject(Store);

    boardId = input.required<number>();
    board$: Observable<BoardModel>;
    columns$: Observable<ColumnModel[]>;

    ngOnInit() {
        this.store.dispatch(actionsColumns.loadByBoardId({ board_id: this.boardId() }));
        this.board$ = this.store
            .select(selectBoard.byId(this.boardId()))
            .pipe(filter((board) => board !== null));
        this.columns$ = this.store.select(selectStatuses.byListId(this.boardId()));
    }
}
