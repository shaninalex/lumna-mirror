import { Component, computed, inject, Input } from '@angular/core';
import { BoardStore } from '@entities/board';
import { BoardCreateFeature } from '@features/board-create';
import { RouterLink } from '@angular/router';

@Component({
    selector: 'app-boards-list',
    imports: [BoardCreateFeature, RouterLink],
    template: `
        <div class="bg-gray-200 card">
            <h2 class="font-medium text-lg mb-4">Boards</h2>
            @if (boards(); as boards) {
                <nav class="flex flex-wrap gap-4">
                    @for (board of boards; track board.id) {
                        <a [routerLink]="['board', board.id]" class="btn btn-primary">{{
                            board.title
                        }}</a>
                    }
                    <app-board-create-feature [projectId]="projectId" />
                </nav>
            }
        </div>
    `,
})
export class BoardsList {
    @Input() projectId: string;
    private readonly boardStore = inject(BoardStore);
    boards = computed(() => this.boardStore.projectBoards(this.projectId));
}
