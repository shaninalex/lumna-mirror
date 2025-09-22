import {Component} from '@angular/core';
import {RouterOutlet} from '@angular/router';

@Component({
    selector: 'fr-root',
    imports: [
        RouterOutlet,
    ],
    template: `<router-outlet/>`,
})
export class App {}
