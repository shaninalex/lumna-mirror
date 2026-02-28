import {Component, Input} from '@angular/core';

@Component({
    selector: 'app-auth-layout',
    imports: [],
    template: `
        <div class="bg-body-tertiary auth-layout">
            <div>
                @if (hasLogo) {
                    <img src="img/logo-icon.svg" style="width: 64px;" />
                }
                <ng-content/>
            </div>
        </div>
    `,
    styleUrl: './auth.css'
})
export class AuthLayout {
    @Input() hasLogo: boolean;
}
