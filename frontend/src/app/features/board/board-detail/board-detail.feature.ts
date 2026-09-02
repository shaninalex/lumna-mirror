import { AsyncPipe } from '@angular/common';
import type { OnInit } from '@angular/core';
import { Component, inject, input } from '@angular/core';
import { Store } from '@ngrx/store';
import type { ColumnModel } from '@entities/column';
import {
    actionsColumns,
    ColumnItemComponent,
    NewColumnFormComponent,
    selectColumns,
} from '@entities/column';
import { filter, type Observable } from 'rxjs';
import { TimeAgoPipe } from '@shared/utils';
import { selectBoard, type BoardModel } from '@entities/board';
import type { TaskModel } from '@entities/task';
import { actionTask, selectTasks } from '@entities/task';

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
    tasks$: Observable<TaskModel[]>;

    ngOnInit() {
        const _q = { board_id: this.boardId() };
        this.store.dispatch(actionTask.getList({ query: _q }));
        this.store.dispatch(actionsColumns.loadByBoardId(_q));

        this.board$ = this.store
            .select(selectBoard.byId(_q.board_id))
            .pipe(filter((board) => board !== null));

        this.columns$ = this.store.select(selectColumns.byListId(_q.board_id));
        this.tasks$ = this.store.select(selectTasks.byBoardId(_q.board_id));
    }
}
