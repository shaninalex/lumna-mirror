import {Component, inject} from '@angular/core';
import {UiService} from '@client/shared/ui';
import {VerificationFormComponent} from '@client/features/auth';

@Component({
    selector: 'jr-verification',
    imports: [
        VerificationFormComponent
    ],
    template: `
        <jr-verification-form />
    `
})
export class VerificationComponent {
    private ui: UiService = inject(UiService)
    constructor() {
        this.ui.setTitle("Verification");
    }
}
