import { Component } from "@angular/core";
import { AuthLoginFeature } from "@features/auth";
import { RouterLink } from "@angular/router";

@Component({
    selector: "app-login",
    imports: [AuthLoginFeature, RouterLink],
    template: `
        <h3 class="mb-3 font-bold text-lg text-center">Login</h3>

        <div class="mb-3">
            <auth-login-feature />
        </div>

        <a routerLink="/">Forgot password</a>
    `
})
export class LoginPage {}
