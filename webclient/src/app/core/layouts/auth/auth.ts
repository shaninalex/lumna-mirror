import { Component, Input } from '@angular/core';
@Component({
    selector: 'app-auth-layout',
    imports: [],
    template: `
        <div class="w-screen h-screen flex items-center justify-center">
            <div class="overflow-auto w-sm">
                @if (hasLogo) {
                    <img src="img/logo.svg" class="mb-8 w-48 mx-auto" alt="Lumna" />
                }
                <ng-content />
            </div>
        </div>
    `,
})
export class AuthLayout {
    @Input() hasLogo: boolean;
}
