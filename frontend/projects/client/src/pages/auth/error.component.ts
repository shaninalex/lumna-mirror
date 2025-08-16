import {Component, inject} from '@angular/core';
import {UiService} from '@client/shared/ui/services/ui.service';
import {ErrorFormComponent} from '@client/features/auth/components';
import {AuthLayoutComponent} from '../../app/layouts';

@Component({
    selector: 'jr-error',
    imports: [
        ErrorFormComponent,
        AuthLayoutComponent
    ],
    template: `
        <auth-layout title="Error">
            <jr-error-form/>
        </auth-layout>
    `
})
export class ErrorComponent {
    private ui: UiService = inject(UiService)

    constructor() {
        this.ui.setTitle("Error");
    }
}
