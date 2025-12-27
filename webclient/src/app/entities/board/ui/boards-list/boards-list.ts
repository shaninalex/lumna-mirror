import {Component, computed, inject, Input} from '@angular/core';
import {BoardStore} from '@entities/board';
import {BoardCreateFeature} from '@features/board-create';

@Component({
    selector: 'app-boards-list',
    imports: [
        BoardCreateFeature
    ],
    template: `
        <div class="bg-gray-200 card">
            <h2 class="font-medium text-lg mb-4">Boards</h2>
            @if (boards(); as boards) {
                <nav class="flex flex-wrap gap-4">
                    @for (board of boards; track board.id) {
                        <button class="btn btn-primary">{{ board.name }}</button>
                    }
                    <app-board-create-feature [projectId]="projectId" />
                </nav>
            }
        </div>
    `,
})
export class BoardsList {
    @Input() projectId: number
    private readonly boardStore = inject(BoardStore);
    boards = computed(() => this.boardStore.projectBoards(this.projectId));
}
