import {Component, inject} from '@angular/core';
import {ActivatedRoute, Params, RouterLink} from '@angular/router';
import {map, Observable} from 'rxjs';
import {AsyncPipe} from '@angular/common';
import {environment} from '@client/environments/environment.development'
import {LoginForm, RegisterForm} from '@client/features/auth';
import {AuthLayout} from '@client/app/layouts';

@Component({
    selector: 'fr-register',
    imports: [
        AsyncPipe,
        AuthLayout,
        RouterLink,
        RegisterForm
    ],
    template: `
        <fr-auth-layout title="Registration">
            @if (flowID$ | async; as flowID) {
                <fr-registration-form [flowID]="flowID"/>
            }
            Already have an account? <a [routerLink]="['/auth/login']" class="text-blue-600">Login</a>
        </fr-auth-layout>
    `,
})
export class Register {
    activatedRoute = inject(ActivatedRoute)
    flowID$: Observable<string|null> = this.activatedRoute.queryParams.pipe(
        map((params: Params) => {
            if (!params.hasOwnProperty("flow")) {
                window.location.href = environment.AUTH_URL_REGISTRATION_REDIRECT;
                return null;
            }
            return params["flow"];
        })
    )
}
