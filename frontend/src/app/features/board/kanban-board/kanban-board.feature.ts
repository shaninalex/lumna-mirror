import { AsyncPipe } from '@angular/common';
import type { OnInit } from '@angular/core';
import { Component, effect, inject, input } from '@angular/core';
import { Store } from '@ngrx/store';
import type { ColumnModel } from '@entities/column';
import {
    actionsColumns,
    ColumnItemComponent,
    NewColumnFormComponent,
} from '@entities/column';
import { filter, tap, type Observable } from 'rxjs';
import { TimeAgoPipe } from '@shared/utils';
import { selectBoard, type BoardModel } from '@entities/board';
import { TaskCardComponent, actionTask } from '@entities/task';

import type { CdkDragDrop } from '@angular/cdk/drag-drop';
import {
    CdkDrag,
    CdkDragHandle,
    CdkDropList,
    CdkDropListGroup,
} from '@angular/cdk/drag-drop';
import { CdkContextMenuTrigger } from '@angular/cdk/menu';
import type { KanbanCard, KanbanColumn } from './model/kanban.models';
import { KanbanService } from './service';

@Component({
    selector: 'lu-kanban-board-feature',
    imports: [
        CdkDropListGroup,
        CdkDropList,
        CdkDrag,
        CdkDragHandle,
        AsyncPipe,
        NewColumnFormComponent,
        TimeAgoPipe,
        ColumnItemComponent,
        CdkContextMenuTrigger,
        TaskCardComponent,
    ],
    templateUrl: './kanban-board.feature.html',
    providers: [KanbanService],
})
export class KanbanBoardFeature implements OnInit {
    private store = inject(Store);
    private kanban = inject(KanbanService);

    boardId = input.required<number>();
    board$: Observable<BoardModel>;

    columnsWithTasks$!: Observable<(ColumnModel & { tasks: KanbanCard[] })[]>;

    kolumns$: Observable<KanbanColumn[]> = this.kanban.boardData();

    constructor() {
        effect(() => {
            this.kanban.setBoardId(this.boardId());
        });
    }

    ngOnInit() {
        const _q = { board_id: this.boardId() };
        this.store.dispatch(actionTask.getList({ query: _q }));
        this.store.dispatch(actionsColumns.loadByBoardId(_q));
        this.board$ = this.store
            .select(selectBoard.byId(_q.board_id))
            .pipe(filter((board) => board !== null));

        this.kolumns$.pipe(
            tap((data) => console.log(data))
        ).subscribe()
    }

    dropColumn(event: CdkDragDrop<ColumnModel[]>): void {}
    dropTask(event: CdkDragDrop<KanbanCard[]>, column: ColumnModel): void {}
}
