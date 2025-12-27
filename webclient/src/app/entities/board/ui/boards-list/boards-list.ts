import {Component, computed, inject, Input, signal} from '@angular/core';
import {BoardModel, BoardStore} from '@entities/board';

@Component({
    selector: 'app-boards-list',
    imports: [],
    template: `
        <div class="bg-gray-200 card">
            <h2 class="font-medium text-lg mb-4">Boards</h2>
            @if (boards(); as boards) {
                <nav class="flex flex-wrap gap-4">
                    @for (board of boards; track board.id) {
                        <button class="btn btn-primary">{{ board.name }}</button>
                    }
                </nav>
            }
        </div>
    `,
})
export class BoardsList {
    @Input() projectId: number
    private readonly boardStore = inject(BoardStore);
    boards = computed(() => this.boardStore.entities().filter(p => p.id === this.projectId));
}
