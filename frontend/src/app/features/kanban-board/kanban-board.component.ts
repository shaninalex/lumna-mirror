import {Component, inject, Input, OnInit} from '@angular/core';
import {BoardModel} from '@entities/board';
import {actionColumnChangeOrder, ColumnModel, ColumnState} from '@entities/column';
import {NewColumnForm} from '@features/kanban-board/components';
import {ColumnDropdownFeature, ColumnEditNameFeature, TaskInlineFormFeature} from '@features';
import {actionTaskChangeOrder, selectTasksByColumns, TaskModel, TaskState} from '@entities/task';
import {TaskCard} from '@entities/task/ui/task-card/task-card';
import {
    CdkDrag,
    CdkDragDrop,
    CdkDragHandle,
    CdkDropList,
    CdkDropListGroup,
    moveItemInArray,
    transferArrayItem,
} from '@angular/cdk/drag-drop';
import {KanbanApi} from './api/kanban.api';
import {Store} from '@ngrx/store';
import {actionKanbanLoadColumns, KanbanBoardChangeOrderPayload} from './model';
import {selectColumnsByBoardId, selectColumnsById} from '@entities/column/model/column.selectors';
import {combineLatest, map, Observable, switchMap, take} from 'rxjs';
import {AsyncPipe} from '@angular/common';

@Component({
    selector: 'app-kanban-board-feature',
    standalone: true,
    imports: [
        NewColumnForm,
        ColumnEditNameFeature,
        TaskCard,
        CdkDropListGroup,
        CdkDropList,
        CdkDrag,
        CdkDragHandle,
        TaskInlineFormFeature,
        AsyncPipe,
        ColumnDropdownFeature,
    ],
    templateUrl: './kanban-board.component.html',
    styleUrl: './kanban-board.component.css',
    providers: [KanbanApi],
})
export class KanbanBoardFeature implements OnInit {
    private columnStore = inject(Store<ColumnState>);
    private taskStore = inject(Store<TaskState>);
    private kanbanApi = inject(KanbanApi);

    @Input() board!: BoardModel;

    columns$!: Observable<ColumnModel[]>;
    tasks$!: Observable<TaskModel[]>;

    /** View model: columns with tasks */
    columnsWithTasks$!: Observable<(ColumnModel & { tasks: TaskModel[] })[]>;

    ngOnInit(): void {
        this.columnStore.dispatch(actionKanbanLoadColumns({boardId: this.board.id}));
        this.columns$ = this.columnStore.select(selectColumnsByBoardId(this.board.id));
        this.tasks$ = this.columns$.pipe(
            switchMap((columns) =>
                this.taskStore.select(selectTasksByColumns(columns.map((c) => c.id))),
            ),
        );
        this.columnsWithTasks$ = combineLatest([this.columns$, this.tasks$]).pipe(
            map(([columns, tasks]) =>
                columns.map((col) => ({
                    ...col,
                    tasks: tasks
                        .filter((t) => t.column_id === col.id)
                        .sort((a, b) => a.order - b.order),
                })),
            ),
        );
    }

    // Drag & Drop – Tasks
    dropTask(event: CdkDragDrop<TaskModel[]>, column: ColumnModel): void {
        const board = this.board;
        if (!board) return;

        const isSameList = event.previousContainer === event.container;

        if (isSameList) {
            // move task in a single column ( just change the order )
            moveItemInArray(event.container.data, event.previousIndex, event.currentIndex);
            const payload: KanbanBoardChangeOrderPayload = {
                moveType: 'task',
                columnId: column.id,
                tasks: this._buildTasksPayload(event.container.data)
            };
            const activityLog = this._logChangeTaskOrder(event.previousIndex, event.currentIndex, event.item.data.id)
            if (activityLog) payload.activity = activityLog
            this.kanbanApi
                .Patch(board.id, payload)
                .subscribe(() => this.taskStore.dispatch(actionTaskChangeOrder(payload)));
        } else {
            // move task between columns and change order
            const activityLog = this._logMovingTaskBetweenColumns(event.item.data.column_id, column.id, event.item.data.id)
            transferArrayItem(
                event.previousContainer.data,
                event.container.data,
                event.previousIndex,
                event.currentIndex,
            );
            const payload: KanbanBoardChangeOrderPayload = {
                moveType: 'task',
                columns: [
                    {
                        id: event.item.data.column_id,
                        tasks: this._buildTasksPayload(event.previousContainer.data),
                    },
                    {
                        id: column.id,
                        tasks: this._buildTasksPayload(event.container.data),
                    },
                ],
            };
            if (activityLog) payload.activity = activityLog
            this.kanbanApi
                .Patch(board.id, payload)
                .subscribe(() => this.taskStore.dispatch(actionTaskChangeOrder(payload)));
        }
    }

    // Drag & Drop – Columns
    dropColumn(event: CdkDragDrop<ColumnModel[]>): void {
        this.columns$.pipe(take(1)).subscribe((columns) => {
            const reordered = [...columns];

            moveItemInArray(reordered, event.previousIndex, event.currentIndex);

            const updatedColumns = reordered.map((col, index) => ({
                id: col.id,
                order: index,
            }));

            const payload: KanbanBoardChangeOrderPayload = {
                moveType: 'column',
                columns: updatedColumns,
            };

            const activityLog = this._logChangeColumnOrder(event.previousIndex, event.currentIndex, event.item.data.id)
            if (activityLog) payload.activity = activityLog

            this.kanbanApi
                .Patch(this.board.id, payload)
                .subscribe(() =>
                    this.columnStore.dispatch(actionColumnChangeOrder({columns: updatedColumns})),
                );
        });
    }

    private _buildTasksPayload(tasks: TaskModel[]) {
        return tasks.map((task, index) => ({
            id: task.id,
            order: index,
        }));
    }

    private _getColumnById(columnId: number): ColumnModel | undefined {
        return this.columnStore.selectSignal(selectColumnsById(columnId))()
    }

    private _logMovingTaskBetweenColumns(fromColumnId: number, toColumnId: number, taskId: number): any {
        const fromColumn = this._getColumnById(fromColumnId)
        const toColumn = this._getColumnById(toColumnId)
        if (!fromColumn || !toColumn) return null
        return {
            entity_type: "task",
            entity_id: taskId,
            summary: `Move from "${fromColumn.title}" to "${toColumn.title}"`,
            action: "update",
        }
    }

    private _logChangeColumnOrder(previousIndex: number, currentIndex: number, columnId: number): any {
        if (previousIndex === currentIndex) return null
        return {
            entity_type: "column",
            entity_id: columnId,
            summary: `Reorder from "${previousIndex}" to "${currentIndex}"`,
            action: "reorder",
        }
    }

    private _logChangeTaskOrder(previousIndex: number, currentIndex: number, taskId: number): any {
        if (previousIndex === currentIndex) return null
        return {
            entity_type: "task",
            entity_id: taskId,
            summary: `Reorder from "${previousIndex}" to "${currentIndex}"`,
            action: "reorder",
        }
    }
}
