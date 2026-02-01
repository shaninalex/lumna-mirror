import { Component, inject, Input, OnInit } from '@angular/core';
import { BoardModel } from '@entities/board';
import { actionColumnChangeOrder, ColumnModel, ColumnState } from '@entities/column';
import { NewColumnForm } from '@features/kanban-board/components';
import { CdkMenu, CdkMenuTrigger } from '@angular/cdk/menu';
import {
    ColumnDeleteFeature,
    ColumnEditNameFeature,
    TaskFastFormFeature,
} from '@root/src/app/features';
import { actionTaskChangeOrder, selectTasksByBoardId, TaskModel, TaskState } from '@entities/task';
import { TaskCard } from '@entities/task/ui/task-card/task-card';
import {
    CdkDrag,
    CdkDragDrop,
    CdkDragHandle,
    CdkDropList,
    CdkDropListGroup,
    moveItemInArray,
    transferArrayItem,
} from '@angular/cdk/drag-drop';
import { KanbanApi } from './api/kanban.api';
import { Store } from '@ngrx/store';
import { actionKanbanLoadColumns, KanbanBoardChangeOrderPayload } from './model';
import { selectColumnsByBoardId } from '@entities/column/model/column.selectors';
import { Observable, combineLatest, map, take } from 'rxjs';
import { AsyncPipe } from '@angular/common';

@Component({
    selector: 'app-kanban-board-feature',
    standalone: true,
    imports: [
        NewColumnForm,
        CdkMenu,
        CdkMenuTrigger,
        ColumnDeleteFeature,
        ColumnEditNameFeature,
        TaskCard,
        CdkDropListGroup,
        CdkDropList,
        CdkDrag,
        CdkDragHandle,
        TaskFastFormFeature,
        AsyncPipe,
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
        this.columnStore.dispatch(actionKanbanLoadColumns({ boardId: this.board.id }));

        this.columns$ = this.columnStore.select(selectColumnsByBoardId(this.board.id));

        this.tasks$ = this.taskStore.select(selectTasksByBoardId(this.board.id));

        this.columnsWithTasks$ = combineLatest([this.columns$, this.tasks$]).pipe(
            map(([columns, tasks]) =>
                columns.map((col) => ({
                    ...col,
                    tasks: tasks.filter((t) => t.column_id === col.id),
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
            moveItemInArray(event.container.data, event.previousIndex, event.currentIndex);

            const payload: KanbanBoardChangeOrderPayload = {
                moveType: 'task',
                columnId: column.id,
                tasks: this._buildTasksPayload(event.container.data),
            };

            this.kanbanApi
                .Patch(board.id, payload)
                .subscribe(() => this.taskStore.dispatch(actionTaskChangeOrder(payload)));
        } else {
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

            const updatedColumns = reordered.map((col) => ({
                id: col.id,
                order: col.order,
            }));

            const payload: KanbanBoardChangeOrderPayload = {
                moveType: 'column',
                columns: updatedColumns,
            };

            this.kanbanApi
                .Patch(this.board.id, payload)
                .subscribe(() =>
                    this.columnStore.dispatch(actionColumnChangeOrder({ columns: updatedColumns })),
                );
        });
    }

    private _buildTasksPayload(tasks: TaskModel[]) {
        return tasks.map((task, index) => ({
            id: task.id,
            order: index,
        }));
    }
}
