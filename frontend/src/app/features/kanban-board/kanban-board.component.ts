import { Component, inject, Input, OnInit, signal } from '@angular/core';
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
import { actionKanbanLoadColumns } from './model';
import { selectColumnsByBoardId } from '@entities/column/model/column.selectors';

@Component({
    selector: 'app-kanban-board-feature',
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
    ],
    templateUrl: './kanban-board.component.html',
    providers: [KanbanApi],
})
export class KanbanBoardFeature implements OnInit {
    private columnStore = inject(Store<ColumnState>);
    private taskStore = inject(Store<TaskState>);
    private kanbanApi = inject(KanbanApi);
    @Input() board: BoardModel;
    column_length: number;

    columns = signal<ColumnModel[]>([]);
    tasks = signal<TaskModel[]>([]);

    ngOnInit() {
        this.columnStore.dispatch(actionKanbanLoadColumns({ boardId: this.board.id }));
        this.columnStore
            .select(selectColumnsByBoardId(this.board.id))
            .subscribe((columns) => this.columns.set(columns));

        this.taskStore.select(selectTasksByBoardId(this.board.id)).subscribe((tasks) => {
            this.tasks.set(tasks);
        });
    }

    listTasks(column_id: string): TaskModel[] {
        return this.tasks().filter((t) => t.column_id === column_id);
    }

    dropTask(event: CdkDragDrop<TaskModel[]>, l: ColumnModel) {
        const isSameList = event.previousContainer === event.container;
        const board = this.board;
        if (!board) return;

        if (isSameList) {
            moveItemInArray(event.container.data, event.previousIndex, event.currentIndex);
            const p = {
                listId: l.id,
                tasks: this._buildTasksPayload(event.container.data),
            };
            this.kanbanApi
                .Patch(board.id, p)
                .subscribe(() => this.taskStore.dispatch(actionTaskChangeOrder(p)));
        } else {
            transferArrayItem(
                event.previousContainer.data,
                event.container.data,
                event.previousIndex,
                event.currentIndex,
            );
            this._updateTaskOrders(event.previousContainer.data);
            this._updateTaskOrders(event.container.data);

            const p = {
                columns: [
                    {
                        id: event.item.data.list_id,
                        tasks: this._buildTasksPayload(event.previousContainer.data),
                    },
                    { id: l.id, tasks: this._buildTasksPayload(event.container.data) },
                ],
            };
            this.kanbanApi
                .Patch(board.id, p)
                .subscribe(() => this.taskStore.dispatch(actionTaskChangeOrder(p)));
        }
    }

    dropList(event: CdkDragDrop<ColumnModel[]>) {
        moveItemInArray(this.columns(), event.previousIndex, event.currentIndex);
        this.columns().forEach((list, index) => (list.order = index));

        const board = this.board;
        if (!board) return;

        const p = {
            columns: this.columns().map((list) => ({ id: list.id, order: list.order })),
        };
        this.kanbanApi
            .Patch(board.id, p)
            .subscribe(() => this.columnStore.dispatch(actionColumnChangeOrder(p)));
    }

    private _updateTaskOrders(tasks: TaskModel[]): void {
        tasks.forEach((task, index) => (task.order = index));
    }

    private _buildTasksPayload(tasks: TaskModel[]): { id: string; order: number }[] {
        return tasks.map((task) => ({ id: task.id, order: task.order }));
    }
}
