import { Component, inject, OnInit } from '@angular/core';
import { AuthLoginFeature } from '@features/index';
import { UiService } from '@shared/ui';

@Component({
    selector: 'app-login',
    imports: [AuthLoginFeature],
    template: `
        <auth-login-feature />

        <hr class="my-4 border-gray-300" />
        <div class="text-center">
            <a routerLink="#" class="text-gray-500 underline">Restore</a>
        </div>
    `,
})
export class Login implements OnInit {
    private ui = inject(UiService);

    ngOnInit(): void {
        this.ui.setPageTitle('Login');
    }
}
