import {Component} from '@angular/core';
import {RouterLink} from '@angular/router';

@Component({
    selector: 'fr-session-expired.component',
    imports: [
        RouterLink
    ],
    template: `
        <p>Your session has been expired. Please login</p>
        <a [routerLink]="['auth', 'login']">login</a>
    `,
})
export class SessionExpiredComponent {

}
