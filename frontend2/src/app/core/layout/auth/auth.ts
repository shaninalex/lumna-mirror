import { Component, Input } from "@angular/core";

@Component({
    selector: "app-auth-layout",
    imports: [],
    template: `
        <div class="h-screen flex items-center justify-center">
            <div>
                <img src="img/logo-h.svg" class="w-64 block mb-4" />
                <ng-content />
            </div>
        </div>
    `
})
export class AuthLayout {}
