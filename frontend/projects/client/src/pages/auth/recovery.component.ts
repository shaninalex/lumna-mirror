import {Component, inject} from '@angular/core';
import {UiService} from '@client/shared/ui';
import {RecoveryFormComponent} from '@client/features/auth';

@Component({
    selector: 'jr-recovery',
    imports: [
        RecoveryFormComponent
    ],
    template: `<jr-recovery-form />`
})
export class RecoveryComponent {
    private ui: UiService = inject(UiService)

    constructor() {
        this.ui.setTitle("Recovery");
    }
}
