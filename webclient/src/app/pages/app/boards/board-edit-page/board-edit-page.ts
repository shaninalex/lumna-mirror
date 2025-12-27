import {Component, inject, signal} from '@angular/core';
import {ActivatedRoute} from '@angular/router';
import {BoardModel} from '@entities/board';
import {BoardDeleteFeature, BoardEditFeature} from '@root/src/app/features';

@Component({
    selector: 'app-board-edit-page',
    imports: [
        BoardEditFeature,
        BoardDeleteFeature
    ],
    template: `
        <h1>Board "{{ board()?.name }}" Edit</h1>
        <app-board-edit-feature [board]="board()"/>
        <div class="mb-4"></div>
        <app-board-delete-feature [board]="board()"/>
    `,
})
export class BoardEditPage {
    route = inject(ActivatedRoute);
    board = signal<BoardModel | undefined>(undefined);

    constructor() {
        this.route.data.subscribe(data => {
            this.board.set(data['board'])
        })
    }
}
