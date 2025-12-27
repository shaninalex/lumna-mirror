import {Component, input} from '@angular/core';
import {BoardModel} from '@entities/board';

@Component({
    selector: 'app-board-delete-feature',
    imports: [],
    template: `
        <p>
            board-delete works!
        </p>
    `,
})
export class BoardDeleteFeature {
    board = input<BoardModel>();

}
