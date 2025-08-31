import {Component, inject} from '@angular/core';
import {ActivatedRoute, Params, RouterLink} from '@angular/router';
import {map, Observable} from 'rxjs';
import {AsyncPipe} from '@angular/common';
import {AuthLayout} from '@client/app/layouts';
import {VerificationForm} from '@client/features/auth/verification-form';

@Component({
    selector: 'fr-verification',
    imports: [
        AsyncPipe,
        AuthLayout,
        VerificationForm
    ],
    template: `
        <fr-auth-layout title="Verification">
            @if (flowID$ | async; as flowID) {
                <fr-verification-form [flowID]="flowID"/>
            }
        </fr-auth-layout>
    `,
})
export class Verification {
    activatedRoute = inject(ActivatedRoute)
    flowID$: Observable<string|null> = this.activatedRoute.queryParams.pipe(
        map((params: Params) => {
            // if (!params.hasOwnProperty("flow")) {
            //     window.location.href = environment.AUTH_URL_REGISTRATION_REDIRECT;
            //     return null;
            // }
            return params["flow"];
        })
    )
}
