import {Component, inject} from '@angular/core';
import {ActivatedRoute} from '@angular/router';
import {AuthLayout} from '@client/app/layouts';
import {KratosFormRenderer} from '@dev/ui/kratos';

@Component({
    selector: 'fr-verification',
    imports: [
        AuthLayout,
        KratosFormRenderer,
    ],
    template: `
        <fr-auth-layout title="Verification" [ready]="!!verificationForm">
            <ui-form-renderer [flow]="verificationForm" />
        </fr-auth-layout>
    `,
})
export class Verification {
    activatedRoute = inject(ActivatedRoute)
    verificationForm = this.activatedRoute.snapshot.data['verificationForm']
}
