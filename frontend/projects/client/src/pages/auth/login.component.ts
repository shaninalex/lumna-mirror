import {Component, inject} from '@angular/core';
import {RouterLink} from '@angular/router';
import {UiService} from '@client/shared/ui';
import {LoginFormComponent} from '@client/features/auth';

@Component({
    selector: 'jr-login',
    imports: [
        RouterLink,
        LoginFormComponent,
    ],
    template: `
        <jr-login-form />
        Don't have an account? <a [routerLink]="['/auth/registration']" class="underline">Registration</a> <br />
        Forgot password? <a [routerLink]="['/auth/recovery']" class="underline">Recovery</a>
    `
})
export class LoginComponent {
    private ui: UiService = inject(UiService);
    constructor() {
        this.ui.setTitle("Login");
    }
}
