import {Component, computed, inject, input} from '@angular/core';
import {BoardModel} from '@entities/board';
import {ListStore} from '@entities/list';
import {NewColumnForm} from '@features/kanban-board/components';
import {CdkMenu, CdkMenuTrigger} from '@angular/cdk/menu';
import {ListDeleteFeature} from '@root/src/app/features';

@Component({
    selector: 'app-kanban-board-feature',
    imports: [
        NewColumnForm,
        CdkMenu,
        CdkMenuTrigger,
        ListDeleteFeature
    ],
    template: `
        <div class="flex items-start gap-4 overflow-x-scroll w-full">
            @for (l of lists(); track l.id) {
                <div class="card bg-gray-100 w-[280px] flex-shrink-0">
                    <div class="flex items-center justify-between">
                        <div class="text-lg font-medium">{{ l.name }}</div>
                        <div>
                            <button [cdkMenuTriggerFor]="menu"  class="cursor-pointer">
                                <i class="fa-solid fa-ellipsis"></i>
                            </button>
                            <ng-template #menu>
                                <div class="bg-white border border-gray-200 rounded-xl p-4" cdkMenu>
                                    <app-list-delete-feature [listId]="l.id" [listName]="l.name" />
                                </div>
                            </ng-template>
                        </div>
                    </div>
                </div>
            }
            <app-new-column-form [board]="board()" />
        </div>
    `,
})
export class KanbanBoardFeature {
    private readonly listStore = inject(ListStore);

    board = input<BoardModel>()
    lists = computed(() => {
        const b = this.board()
        if (b) {
            return this.listStore.boardLists(b.id)
        }
        return []
    });
}
