import {Component, inject} from '@angular/core';
import {ActivatedRoute, Params, RouterLink} from '@angular/router';
import {AuthLayout} from '@client/app/layouts';
import {RegistrationFlow, VerificationFlow} from '@ory/kratos-client';
import {KratosFormRenderer} from '@dev/ui/kratos';

@Component({
    selector: 'fr-verification',
    imports: [
        AuthLayout,
        KratosFormRenderer,
    ],
    template: `
        @if (verificationForm) {
            <fr-auth-layout title="Verification">
                <ui-form-renderer [flow]="verificationForm" />
            </fr-auth-layout>
        }
    `,
})
export class Verification {
    activatedRoute = inject(ActivatedRoute)
    verificationForm: VerificationFlow | undefined = this.activatedRoute.snapshot.data['verificationForm']
}
