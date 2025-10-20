import {Component} from '@angular/core';
import {RouterOutlet} from '@angular/router';

@Component({
    selector: 'lu-projects-root',
    imports: [
        RouterOutlet
    ],
    template: `
        <router-outlet/>`
})
export class ProjectsRootPageComponent {
}
