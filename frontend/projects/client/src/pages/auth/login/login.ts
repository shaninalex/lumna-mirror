import {Component, inject} from '@angular/core';
import {ActivatedRoute, Params, RouterLink} from '@angular/router';
import {map, Observable} from 'rxjs';
import {AsyncPipe} from '@angular/common';
import {environment} from '@client/environments/environment.development'
import {LoginForm} from '@client/features/auth';
import {AuthLayout} from '@client/app/layouts';

@Component({
    selector: 'fr-login',
    imports: [
        AsyncPipe,
        LoginForm,
        AuthLayout,
        RouterLink
    ],
    template: `
        <fr-auth-layout title="Login">
            @if (flowID$ | async; as flowID) {
                <fr-login-form [flowID]="flowID"/>
            }
            Don't have an account? <a [routerLink]="['/auth/registration']" class="underline">Registration</a> <br />
            Forgot password? <a [routerLink]="['/auth/recovery']" class="underline">Recovery</a>
        </fr-auth-layout>
    `,
})
export class Login {
    activatedRoute = inject(ActivatedRoute)
    flowID$: Observable<string|null> = this.activatedRoute.queryParams.pipe(
        map((params: Params) => {
            if (!params.hasOwnProperty("flow")) {
                window.location.href = environment.AUTH_URL_LOGIN_REDIRECT;
                return null;
            }
            return params["flow"];
        })
    )
}
