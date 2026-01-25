import { Component, inject, signal } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { BoardModel } from '@entities/board';
import { BoardDeleteFeature, BoardEditFeature } from '@root/src/app/features';

@Component({
    selector: 'app-board-edit-page',
    imports: [BoardEditFeature, BoardDeleteFeature],
    template: `
        <h1>Board "{{ board()?.title }}" Edit</h1>
        <app-board-edit-feature [board]="board()" />
        <div class="my-4 border-b border-gray-200"></div>
        <app-board-delete-feature [board]="board()" [projectId]="projectId" />
    `,
})
export class BoardEditPage {
    route = inject(ActivatedRoute);
    board = signal<BoardModel | undefined>(undefined);
    projectId: string;

    constructor() {
        this.route.data.subscribe((data) => {
            this.board.set(data['board']);
        });
        this.route.params.subscribe((params) => {
            this.projectId = params['id'];
        });
    }
}
