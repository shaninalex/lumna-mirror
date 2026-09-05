import { AsyncPipe } from '@angular/common';
import type { OnInit } from '@angular/core';
import { Component, effect, inject, input } from '@angular/core';
import { Store } from '@ngrx/store';
import { actionsColumns, NewColumnFormComponent } from '@entities/column';
import { filter, type Observable } from 'rxjs';
import { TimeAgoPipe } from '@shared/utils';
import { selectBoard, type BoardModel } from '@entities/board';
import { TaskCardComponent, actionTask, TaskInlineForm } from '@entities/task';
import type { CdkDragDrop } from '@angular/cdk/drag-drop';
import {
    CdkDrag,
    CdkDragHandle,
    CdkDropList,
    CdkDropListGroup,
    moveItemInArray,
    transferArrayItem,
} from '@angular/cdk/drag-drop';
import type { KanbanCard, KanbanColumn } from './model/kanban.models';
import { KanbanService } from './service';
import { actionKanban } from './model';

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
        TaskCardComponent,
        TaskInlineForm,
    ],
    templateUrl: './kanban-board.feature.html',
    styleUrl: './kanban-board.feature.css',
    providers: [KanbanService],
})
export class KanbanBoardFeature implements OnInit {
    private store = inject(Store);
    private kanban = inject(KanbanService);

    boardId = input.required<number>();
    board$: Observable<BoardModel>;
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
    }

    dropColumn(event: CdkDragDrop<KanbanColumn[]>): void {
        const data = this.kanban.getData();
        moveItemInArray(data, event.previousIndex, event.currentIndex);
        this.store.dispatch(
            actionKanban.dropColumn({
                event: {
                    id: event.item.data.id,
                    previous_ndex: event.previousIndex,
                    current_index: event.currentIndex,
                    board_id: this.boardId(),
                    columns_order: data.map((c) => c.id),
                },
            }),
        );
    }

    dropTask(event: CdkDragDrop<KanbanCard[]>, column: KanbanColumn): void {
        const isSameList = event.previousContainer === event.container;

        if (isSameList) {
            moveItemInArray(event.container.data, event.previousIndex, event.currentIndex);
            this.store.dispatch(
                actionKanban.moveTask({
                    event: {
                        board_id: this.boardId(),
                        column_id: column.id,
                        tasks: event.container.data.map((t) => t.id),
                    },
                }),
            );
        } else {
            transferArrayItem(
                event.previousContainer.data,
                event.container.data,
                event.previousIndex,
                event.currentIndex,
            );
            const data = {
                event: {
                    board_id: this.boardId(),
                    from: {
                        column_id: event.item.data.column,
                        tasks: event.previousContainer.data.map((t) => t.id),
                    },
                    to: {
                        column_id: column.id,
                        tasks: event.container.data.map((t) => t.id),
                    },
                },
            };
            this.store.dispatch(actionKanban.transferTask(data));
        }
    }

    public columnsAmount(): number {
        return this.kanban.getColumnsLength()
    }
}
