import {Component, inject} from '@angular/core';
import {ActivatedRoute, RouterLink} from '@angular/router';
import {AuthLayout} from '@client/shared/layouts/auth-layout/auth-layout';
import {AuthRegistrationFeature} from '@client/features/auth/auth-registration-feature';

@Component({
    selector: 'kr-registration',
    imports: [
        RouterLink,
        AuthRegistrationFeature,
        AuthLayout,
    ],
    template: `
        <fr-auth-layout title="Registration">
            <kr-auth-registration-feature [form]="form" />
            Already have an account? <a [routerLink]="['/auth/login']" class="underline">Login</a>
        </fr-auth-layout>
    `
})
export class Registration {
    form = inject(ActivatedRoute).snapshot.data['form']
}
