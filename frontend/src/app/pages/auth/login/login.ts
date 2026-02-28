import { Component, inject, OnInit } from '@angular/core';
import { AuthLoginFeature } from '@features/index';
import { UiService } from '@shared/ui';
import {RouterLink} from '@angular/router';

@Component({
    selector: 'app-login',
    imports: [AuthLoginFeature, RouterLink],
    template: `
        <h1 class="title">
            Login
        </h1>
        <div class="box">
            <auth-login-feature />
        </div>
        <p class="has-text-centered">
            <a routerLink="/">Forgot password</a>
        </p>
    `,
})
export class Login implements OnInit {
    private ui = inject(UiService);

    ngOnInit(): void {
        this.ui.setPageTitle('Login');
    }
}
