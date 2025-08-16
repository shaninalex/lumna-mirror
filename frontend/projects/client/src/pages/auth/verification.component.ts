import {Component, inject} from '@angular/core';
import {UiService} from '@client/shared/ui';
import {VerificationFormComponent} from '@client/features/auth';
import {AuthLayoutComponent} from '../../app/layouts';

@Component({
    selector: 'jr-verification',
    imports: [
        VerificationFormComponent,
        AuthLayoutComponent
    ],
    template: `
        <auth-layout title="Verification">
            <jr-verification-form />
        </auth-layout>
    `
})
export class VerificationComponent {
    private ui: UiService = inject(UiService)
    constructor() {
        this.ui.setTitle("Verification");
    }
}
