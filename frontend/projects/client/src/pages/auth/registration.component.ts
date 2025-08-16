import {Component, inject} from '@angular/core';
import {RouterLink} from '@angular/router';
import {UiService} from '@client/shared/ui';
import {RegistrationFormComponent} from '@client/features/auth';

@Component({
    selector: 'jr-registration',
    imports: [
        RegistrationFormComponent,
        RouterLink
    ],
    template: `
        <jr-registration-form />
        Already have an account? <a [routerLink]="['/auth/login']" class="text-blue-600">Login</a>
    `
})
export class RegistrationComponent {
    private ui: UiService = inject(UiService)

    constructor() {
        this.ui.setTitle("Registration");
    }
}
