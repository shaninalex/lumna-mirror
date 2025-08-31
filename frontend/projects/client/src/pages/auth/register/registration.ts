import {Component, inject} from '@angular/core';
import {ActivatedRoute, RouterLink} from '@angular/router';
import {AuthLayout} from '@client/app/layouts';
import {KratosFormRenderer} from '@dev/ui/kratos';

@Component({
    selector: 'fr-register',
    imports: [
        AuthLayout,
        RouterLink,
        KratosFormRenderer,
    ],
    template: `
        <fr-auth-layout title="Registration" [ready]="!!registrationForm">
            <ui-form-renderer [flow]="registrationForm"/>
            Already have an account? <a [routerLink]="['/auth/login']" class="text-blue-600">Login</a>
        </fr-auth-layout>
    `,
})
export class Register {
    activatedRoute = inject(ActivatedRoute)
    registrationForm = this.activatedRoute.snapshot.data['registrationForm']
}
