import {Component, inject} from '@angular/core';
import {UiService} from '@client/shared/ui';
import {RecoveryFormComponent} from '@client/features/auth';
import {AuthLayoutComponent} from '../../app/layouts';

@Component({
    selector: 'jr-recovery',
    imports: [
        RecoveryFormComponent,
        AuthLayoutComponent
    ],
    template: `
        <auth-layout title="Recovery">
            <jr-recovery-form/>
        </auth-layout>
    `
})
export class RecoveryComponent {
    private ui: UiService = inject(UiService)

    constructor() {
        this.ui.setTitle("Recovery");
    }
}
