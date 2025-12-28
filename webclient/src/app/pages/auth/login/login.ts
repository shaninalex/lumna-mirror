import { Component } from '@angular/core';
import { AuthLayout } from '@core/layouts'
import {RouterLink} from '@angular/router';
import {LoginFormFeature} from '@features/auth';

@Component({
    selector: 'app-login',
    imports: [AuthLayout, RouterLink, LoginFormFeature],
    template: `
        <app-auth-layout [hasLogo]="true">
            <app-login-form-feature />
            <hr class="my-4 border-gray-300">
            <div class=" text-center">
                <a routerLink="/auth/restore" class="text-gray-500 underline">Restore login</a>
            </div>
        </app-auth-layout>
    `,
    styleUrl: './login.css',
})
export class Login {

}
