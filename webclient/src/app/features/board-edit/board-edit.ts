import {Component, input} from '@angular/core';
import {BoardModel} from '@entities/board';

@Component({
    selector: 'app-board-edit-feature',
    imports: [],
    template: `
        <p>
            board-edit works!
        </p>
    `,
})
export class BoardEditFeature {
    board = input<BoardModel>();
}
