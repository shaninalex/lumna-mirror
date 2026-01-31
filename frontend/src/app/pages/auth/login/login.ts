import { Component } from '@angular/core';
import { AuthLayout } from '@core/layouts';
import { AuthLoginFeature } from '@features/index';

@Component({
    selector: 'app-login',
    imports: [AuthLoginFeature, AuthLayout],
    template: `
        <app-auth-layout [hasLogo]="true">
            <auth-login-feature />

            <hr class="my-4 border-gray-300" />
            <div class="text-center">
                <a routerLink="#" class="text-gray-500 underline">Restore</a>
            </div>
        </app-auth-layout>
    `,
})
export class Login {}
