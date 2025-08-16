import {Component, inject} from '@angular/core';
import {UiService} from '@client/shared/ui/services/ui.service';
import {ErrorFormComponent} from '@client/features/auth/components';

@Component({
    selector: 'jr-error',
    imports: [
        ErrorFormComponent
    ],
    template: `<jr-error-form />`
})
export class ErrorComponent {
    private ui: UiService = inject(UiService)

    constructor() {
        this.ui.setTitle("Error");
    }
}
