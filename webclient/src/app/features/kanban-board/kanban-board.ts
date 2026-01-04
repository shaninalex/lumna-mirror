import {Component, effect, inject, input, signal} from '@angular/core';
import { BoardModel } from '@entities/board';
import {ListModel, ListStore} from '@entities/list';
import { NewColumnForm } from '@features/kanban-board/components';
import { CdkMenu, CdkMenuTrigger } from '@angular/cdk/menu';
import {ListDeleteFeature, ListEditNameFeature, TaskFastFormFeature} from '@root/src/app/features';
import { TaskModel, TaskStore } from '@entities/task';
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

@Component({
    selector: 'app-kanban-board-feature',
    imports: [
        NewColumnForm,
        CdkMenu,
        CdkMenuTrigger,
        ListDeleteFeature,
        ListEditNameFeature,
        TaskCard,
        CdkDropListGroup,
        CdkDropList,
        CdkDrag,
        CdkDragHandle,
        TaskFastFormFeature,
    ],
    template: `
        <div cdkDropList
            cdkDropListOrientation="horizontal"
            [cdkDropListData]="lists()"
            (cdkDropListDropped)="dropList($event)">

            <div cdkDropListGroup
                class="flex items-start gap-4 overflow-x-scroll w-full">
                @for (l of lists(); track l.id) {
                    <div cdkDrag class="card bg-gray-100 w-[280px] flex-shrink-0">
                        <div class="flex items-center justify-between gap-2 mb-4">
                            <app-list-edit-name-feature [list]="l" />
                            <div>
                                <button cdkDragHandle class="cursor-pointer text-gray-400">
                                    <i class="fa-regular fa-hand"></i>
                                </button>
                                <button [cdkMenuTriggerFor]="menu" class="cursor-pointer">
                                    <i class="fa-solid fa-ellipsis"></i>
                                </button>
                                <ng-template #menu>
                                    <div class="bg-white border border-gray-200 rounded-xl p-4" cdkMenu>
                                        <app-list-delete-feature [listId]="l.id" [listName]="l.name" />
                                    </div>
                                </ng-template>
                            </div>
                        </div>

                        <div class="min-h-[100px] flex flex-col gap-2 mb-4"
                            cdkDropList
                            [cdkDropListData]="listTasks(l.id)"
                            (cdkDropListDropped)="dropTask($event, l)">
                            @for (task of listTasks(l.id); track task.id) {
                                <div cdkDrag [cdkDragData]="task">
                                    <app-task-card [task]="task" />
                                </div>
                            }
                        </div>

                        <app-task-fast-form-feature [list]="l" [task_count]="listTasks(l.id).length" />
                    </div>
                }
                <app-new-column-form [board]="board()" [lists_length]="lists_length" />
            </div>
        </div>
    `,
    providers: [KanbanApi]
})
export class KanbanBoardFeature {
    private readonly listStore = inject(ListStore);
    private readonly taskStore = inject(TaskStore);
    private service = inject(KanbanApi);

    lists_length: number;
    lists = signal<ListModel[]>([]);
    tasks = signal<TaskModel[]>([]);

    constructor() {
        effect(() => {
            const board = this.board();
            if (!board) return;
            this.lists.set(this.listStore.boardLists(board.id));
            this.tasks.set(this.taskStore.boardTasks(board.id));
        })
    }

    board = input<BoardModel>();

    listTasks(list_id: number): TaskModel[] {
        return this.tasks().filter(t => t.list_id === list_id)
    }

    dropTask(event: CdkDragDrop<TaskModel[]>, l: ListModel) {
        const isSameList = event.previousContainer === event.container;
        const board = this.board();
        if (!board) return;

        if (isSameList) {
            moveItemInArray(event.container.data, event.previousIndex, event.currentIndex);
            this._updateTaskOrders(event.previousContainer.data);
            console.log({
                listId: l.id,
                tasks: this._buildTasksPayload(event.container.data)
            });
        } else {
            transferArrayItem(
                event.previousContainer.data,
                event.container.data,
                event.previousIndex,
                event.currentIndex,
            );
            this._updateTaskOrders(event.previousContainer.data);
            this._updateTaskOrders(event.container.data);

            this.service.Patch(board.id, {
                lists: [
                    { id: event.item.data.list_id, tasks: this._buildTasksPayload(event.previousContainer.data) },
                    { id: l.id, tasks: this._buildTasksPayload(event.container.data) }
                ]
            }).subscribe();
        }
    }

    dropList(event: CdkDragDrop<ListModel[]>) {
        moveItemInArray(this.lists(), event.previousIndex, event.currentIndex);
        this.lists().forEach((list, index) => list.order = index);

        const board = this.board();
        if (!board) return;

        this.service.Patch(board.id, {
            lists: this.lists().map(list => ({ id: list.id, order: list.order }))
        }).subscribe()
    }

    private _updateTaskOrders(tasks: TaskModel[]): void {
        tasks.forEach((task, index) => task.order = index);
    }

    private _buildTasksPayload(tasks: TaskModel[]): { id: number; order: number }[] {
        return tasks.map(task => ({ id: task.id, order: task.order }));
    }
}
