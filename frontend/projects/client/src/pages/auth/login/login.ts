import {Component, inject} from '@angular/core';
import {ActivatedRoute,  RouterLink} from '@angular/router';
import {AuthLayout} from '@client/app/layouts';
import {LoginFlow} from '@ory/kratos-client';
import {KratosFormRenderer} from '@dev/ui/kratos';

@Component({
    selector: 'fr-login',
    imports: [
        AuthLayout,
        RouterLink,
        KratosFormRenderer
    ],
    template: `
        @if (loginFlow) {
            <fr-auth-layout title="Login">
                <ui-form-renderer [flow]="loginFlow"/>
                Don't have an account? <a [routerLink]="['/auth/registration']" class="underline">Registration</a> <br />
                Forgot password? <a [routerLink]="['/auth/recovery']" class="underline">Recovery</a>
            </fr-auth-layout>
        }
    `,
})
export class Login {
    activatedRoute = inject(ActivatedRoute)
    loginFlow: LoginFlow | undefined = this.activatedRoute.snapshot.data['loginForm']
}
