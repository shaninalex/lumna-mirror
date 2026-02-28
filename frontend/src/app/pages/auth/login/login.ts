import { Component, inject, OnInit } from '@angular/core';
import { AuthLoginFeature } from '@features/index';
import { UiService } from '@shared/ui';
import {RouterLink} from '@angular/router';

@Component({
    selector: 'app-login',
    imports: [AuthLoginFeature, RouterLink],
    template: `
        <h3 class="h3 mb-3">Login</h3>

        <div class="mb-3">
            <auth-login-feature />
        </div>

        <a routerLink="/">Forgot password</a>
    `,
})
export class Login implements OnInit {
    private ui = inject(UiService);

    ngOnInit(): void {
        this.ui.setPageTitle('Login');
    }
}
