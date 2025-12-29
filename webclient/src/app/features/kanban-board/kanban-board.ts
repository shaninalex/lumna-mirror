import {Component, computed, inject, input} from '@angular/core';
import {BoardModel} from '@entities/board';
import {ListStore} from '@entities/list';
import {NewColumnForm} from '@features/kanban-board/components';

@Component({
    selector: 'app-kanban-board-feature',
    imports: [
        NewColumnForm
    ],
    template: `
        <div class="flex items-start gap-4 overflow-x-scroll w-full">
            @for (l of lists(); track l.id) {
                <div class="card bg-gray-100 w-[280px] flex-shrink-0">
                    <div class="flex items-center justify-between">
                        <div class="text-lg font-medium">{{ l.name }}</div>
                        <div>
                            <button class="cursor-pointer">
                                <i class="fa-solid fa-ellipsis"></i>
                            </button>
                        </div>
                    </div>
                </div>
            }
            <app-new-column-form [board]="board()" />
        </div>
    `,
})
export class KanbanBoardFeature {
    board = input<BoardModel>()
    private readonly listStore = inject(ListStore);
    lists = computed(() => {
        const b = this.board()
        if (b) {
            return this.listStore.boardLists(b.id)
        }
        return []
    });
}
