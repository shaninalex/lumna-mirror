import { Component, inject } from '@angular/core';
import { UiService } from '@shared/ui';

@Component({
    selector: 'lu-login',
    imports: [],
    template: ` <p>login works!</p> `,
})
export class LoginPage {
    private ui = inject(UiService);

    ngOnInit(): void {
        this.ui.setPageTitle("Login")
    }
}
