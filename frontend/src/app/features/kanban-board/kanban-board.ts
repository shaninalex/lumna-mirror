import { Component, effect, inject, input, signal } from '@angular/core';
import { BoardModel } from '@entities/board';
import { ColumnModel, ColumnState } from '@entities/column';
import { NewColumnForm } from '@features/kanban-board/components';
import { CdkMenu, CdkMenuTrigger } from '@angular/cdk/menu';
import {
    ColumnDeleteFeature,
    ColumnEditNameFeature,
    TaskFastFormFeature,
} from '@root/src/app/features';
import { taskEvents, TaskModel, TaskStore } from '@entities/task';
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
// import { listEvents } from '@entities/list/model/list.events';
import { Dispatcher } from '@ngrx/signals/events';

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
    template: `
        <div
            cdkDropList
            cdkDropListOrientation="horizontal"
            [cdkDropListData]="columns()"
            (cdkDropListDropped)="dropList($event)"
        >
            <div cdkDropListGroup class="flex items-start gap-4 overflow-x-scroll w-full">
                @for (c of columns(); track c.id) {
                    <div cdkDrag class="card bg-gray-100 w-[280px] flex-shrink-0">
                        <div class="flex items-center justify-between gap-2 mb-4">
                            <app-column-edit-name-feature [column]="c" />
                            <div>
                                <button cdkDragHandle class="cursor-pointer text-gray-400">
                                    <i class="fa-regular fa-hand"></i>
                                </button>
                                <button [cdkMenuTriggerFor]="menu" class="cursor-pointer">
                                    <i class="fa-solid fa-ellipsis"></i>
                                </button>
                                <ng-template #menu>
                                    <div
                                        class="bg-white border border-gray-200 rounded-xl p-4"
                                        cdkMenu
                                    >
                                        <app-column-delete-feature
                                            [columnId]="c.id"
                                            [columnName]="c.title"
                                        />
                                    </div>
                                </ng-template>
                            </div>
                        </div>

                        <div
                            class="min-h-[100px] flex flex-col gap-2 mb-4"
                            cdkDropList
                            [cdkDropListData]="listTasks(c.id)"
                            (cdkDropListDropped)="dropTask($event, c)"
                        >
                            @for (task of listTasks(c.id); track task.id) {
                                <div cdkDrag [cdkDragData]="task">
                                    <app-task-card [task]="task" />
                                </div>
                            }
                        </div>

                        <app-task-fast-form-feature
                            [list]="c"
                            [task_count]="listTasks(c.id).length"
                        />
                    </div>
                }
                <app-new-column-form [board]="board()" [columns_length]="column_length" />
            </div>
        </div>
    `,
    providers: [KanbanApi],
})
export class KanbanBoardFeature {
    private dispatcher = inject(Dispatcher);
    private listStore = inject(ListStore);
    private taskStore = inject(TaskStore);
    private service = inject(KanbanApi);

    column_length: number;
    columns = signal<ColumnModel[]>([]);
    tasks = signal<TaskModel[]>([]);

    constructor() {
        effect(() => {
            const board = this.board();
            if (!board) return;
            this.columns.set(this.listStore.boardLists(board.id));
            this.tasks.set(this.taskStore.boardTasks(board.id));
        });
    }

    board = input<BoardModel>();

    listTasks(list_id: string): TaskModel[] {
        return this.tasks().filter((t) => t.list_id === list_id);
    }

    dropTask(event: CdkDragDrop<TaskModel[]>, l: ColumnModel) {
        const isSameList = event.previousContainer === event.container;
        const board = this.board();
        if (!board) return;

        if (isSameList) {
            moveItemInArray(event.container.data, event.previousIndex, event.currentIndex);
            const p = {
                listId: l.id,
                tasks: this._buildTasksPayload(event.container.data),
            };
            this.service
                .Patch(board.id, p)
                .subscribe(() => this.dispatcher.dispatch(taskEvents.changeOrder(p)));
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
                lists: [
                    {
                        id: event.item.data.list_id,
                        tasks: this._buildTasksPayload(event.previousContainer.data),
                    },
                    { id: l.id, tasks: this._buildTasksPayload(event.container.data) },
                ],
            };
            this.service
                .Patch(board.id, p)
                .subscribe(() => this.dispatcher.dispatch(taskEvents.changeOrder(p)));
        }
    }

    dropList(event: CdkDragDrop<ColumnModel[]>) {
        moveItemInArray(this.columns(), event.previousIndex, event.currentIndex);
        this.columns().forEach((list, index) => (list.order = index));

        const board = this.board();
        if (!board) return;

        const p = {
            lists: this.columns().map((list) => ({ id: list.id, order: list.order })),
        };
        this.service
            .Patch(board.id, p)
            .subscribe(() => this.dispatcher.dispatch(listEvents.changeOrder(p)));
    }

    private _updateTaskOrders(tasks: TaskModel[]): void {
        tasks.forEach((task, index) => (task.order = index));
    }

    private _buildTasksPayload(tasks: TaskModel[]): { id: string; order: number }[] {
        return tasks.map((task) => ({ id: task.id, order: task.order }));
    }
}
