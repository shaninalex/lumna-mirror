import { RouterOutlet } from '@angular/router';
import { Component } from '@angular/core';

@Component({
    selector: 'app-board-container',
    imports: [RouterOutlet],
    template: ` <router-outlet /> `,
})
export class BoardContainer {}
