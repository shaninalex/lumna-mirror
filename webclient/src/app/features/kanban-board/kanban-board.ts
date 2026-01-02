import { Component, computed, inject, input } from '@angular/core';
import { BoardModel } from '@entities/board';
import { ListModel, ListStore } from '@entities/list';
import { NewColumnForm } from '@features/kanban-board/components';
import { CdkMenu, CdkMenuTrigger } from '@angular/cdk/menu';
import { ListDeleteFeature, ListEditNameFeature } from '@root/src/app/features';
import { KanbanApi } from '@features/kanban-board/api/kanban.api';
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
import { KanbanModel } from './model/kanban.model';

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
    ],
    template: `
        <div cdkDropList
            cdkDropListOrientation="horizontal"
            [cdkDropListData]="kanban()"
            (cdkDropListDropped)="dropList($event)">

            <div cdkDropListGroup 
                class="flex items-start gap-4 overflow-x-scroll w-full">
                @for (l of kanban(); track l.list.id) {
                    <div cdkDrag class="card bg-gray-100 w-[280px] flex-shrink-0">
                        <div class="flex items-center justify-between gap-2 mb-4">
                            <app-list-edit-name-feature [list]="l.list" />
                            <div>
                                <button cdkDragHandle class="cursor-pointer text-gray-400">
                                    <i class="fa-regular fa-hand"></i>
                                </button>
                                <button [cdkMenuTriggerFor]="menu" class="cursor-pointer">
                                    <i class="fa-solid fa-ellipsis"></i>
                                </button>
                                <ng-template #menu>
                                    <div class="bg-white border border-gray-200 rounded-xl p-4" cdkMenu>
                                        <app-list-delete-feature [listId]="l.list.id" [listName]="l.list.name" />
                                    </div>
                                </ng-template>
                            </div>
                        </div>

                        <div cdkDropList
                                [cdkDropListData]="l.tasks"
                                class="min-h-[100px]"
                                (cdkDropListDropped)="dropTask($event)">
                            @for (task of l.tasks; track task.id) {
                                <div cdkDrag>
                                    <app-task-card [task]="task" />
                                </div>
                            }
                        </div>

                    </div>
                }
                <app-new-column-form [board]="board()" />
            </div>
        </div>
    `,
})
export class KanbanBoardFeature {
    private readonly listStore = inject(ListStore);
    private readonly taskStore = inject(TaskStore);

    board = input<BoardModel>();
    kanban = computed(() => {
        const board = this.board();
        if (!board) return null;

        const lists = this.listStore.boardLists(board.id);
        const tasks = this.taskStore.boardTasks(board.id);

        if (!lists.length) return null;

        return lists
            .slice()
            .sort((a, b) => a.order - b.order)
            .map(list => ({
                list: list,
                tasks: tasks
                    .filter(t => t.list_id === list.id)
                    .sort((a, b) => a.order - b.order),
            }));
    });

    dropTask(event: CdkDragDrop<TaskModel[]>) {
        console.log(event)
    }

    dropList(event: CdkDragDrop<any>) {
        console.log(event)
        // moveItemInArray(this.lists, event.previousIndex, event.currentIndex);
        // this.updateListOrders();

        // this.logPayload({
        //     lists: this.lists.map(list => ({ id: list.id, order: list.order }))
        // });
    }
}
