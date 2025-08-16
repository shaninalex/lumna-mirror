import {Component, inject} from '@angular/core';
import {RouterLink} from '@angular/router';
import {UiService} from '@client/shared/ui';
import {LoginFormComponent} from '@client/features/auth';
import {AuthLayoutComponent} from '../../app/layouts';

@Component({
    selector: 'jr-login',
    imports: [
        RouterLink,
        LoginFormComponent,
        AuthLayoutComponent,
    ],
    template: `
        <auth-layout title="Login">
            <jr-login-form/>
            <div>
                <p>Don't have an account? <a [routerLink]="['/auth/registration']" class="underline">Registration</a></p>
                <p>Forgot password? <a [routerLink]="['/auth/recovery']" class="underline">Recovery</a></p>
            </div>
        </auth-layout>
    `
})
export class LoginComponent {
    private ui: UiService = inject(UiService);
    constructor() {
        this.ui.setTitle("Login");
    }
}
