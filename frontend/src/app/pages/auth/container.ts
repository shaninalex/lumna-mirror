import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { AuthLayout } from '@core/layouts';

@Component({
    selector: 'auth-container',
    imports: [RouterOutlet, AuthLayout],
    template: `
        <app-auth-layout [hasLogo]="true">
            <router-outlet />
        </app-auth-layout>
    `,
})
export class AuthContainer {}
