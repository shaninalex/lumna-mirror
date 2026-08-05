import { Component, signal } from '@angular/core';
import { RouterOutlet } from '@angular/router';

@Component({
    selector: 'lu-root',
    imports: [RouterOutlet],
    template: `
        <router-outlet />
    `,
})
export class App {
}
