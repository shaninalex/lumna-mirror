import {Component, inject} from '@angular/core';
import {ActivatedRoute, RouterLink} from '@angular/router';
import {AuthLayout} from '@client/app/layouts';
import {RegistrationFlow} from '@ory/kratos-client';
import {KratosFormRenderer} from '@dev/ui/kratos';

@Component({
    selector: 'fr-register',
    imports: [
        AuthLayout,
        RouterLink,
        KratosFormRenderer
    ],
    template: `
        @if (registrationForm) {
            <fr-auth-layout title="Registration">
                <ui-form-renderer [flow]="registrationForm"/>
                Already have an account? <a [routerLink]="['/auth/login']" class="text-blue-600">Login</a>
            </fr-auth-layout>
        }
    `,
})
export class Register {
    activatedRoute = inject(ActivatedRoute)
    registrationForm: RegistrationFlow | undefined = this.activatedRoute.snapshot.data['registrationForm']
}
