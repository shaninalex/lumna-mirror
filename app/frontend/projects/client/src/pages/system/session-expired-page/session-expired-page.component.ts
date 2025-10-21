import {Component} from '@angular/core';
import {RouterLink} from '@angular/router';

@Component({
    selector: 'lu-session-expired-page',
    imports: [
        RouterLink
    ],
    template: `
        <p>Your session has been expired. Please login</p>
        <a [routerLink]="['auth', 'login']">login</a>
    `,
})
export class SessionExpiredPageComponent {

}
