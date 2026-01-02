import {Component, computed, inject, input} from '@angular/core';
import {BoardModel} from '@entities/board';
import {ListStore} from '@entities/list';
import {NewColumnForm} from '@features/kanban-board/components';
import {CdkMenu, CdkMenuTrigger} from '@angular/cdk/menu';
import {ListDeleteFeature, ListEditNameFeature} from '@root/src/app/features';
import {KanbanApi} from '@features/kanban-board/api/kanban.api';

@Component({
    selector: 'app-kanban-board-feature',
    imports: [
        NewColumnForm,
        CdkMenu,
        CdkMenuTrigger,
        ListDeleteFeature,
        ListEditNameFeature
    ],
    template: `
        <div class="flex items-start gap-4 overflow-x-scroll w-full">
            @for (l of lists(); track l.id) {
                <div class="card bg-gray-100 w-[280px] flex-shrink-0">
                    <div class="flex items-center justify-between gap-2">
                        <app-list-edit-name-feature [list]="l" />
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
    providers: [KanbanApi]
})
export class KanbanBoardFeature {
    private readonly listStore = inject(ListStore);
    private api = inject(KanbanApi);

    board = input<BoardModel>()
    lists = computed(() => {
        const b = this.board()
        if (b) {
            return this.listStore.boardLists(b.id)
        }
        return []
    });
}
