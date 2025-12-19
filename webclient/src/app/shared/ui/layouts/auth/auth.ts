import { Component } from '@angular/core';

@Component({
    selector: 'app-auth-layout',
    imports: [],
    template: `
        <div class="auth-layout">
            <ng-content/>
        </div>
    `,
    styleUrl: './auth.css',
})
export class AuthLayout {

}
