import { Component } from "@angular/core";
import { RouterOutlet } from "@angular/router";
import { AuthLayout } from "@core/layout";

@Component({
    selector: "app-auth-wrapper",
    imports: [RouterOutlet, AuthLayout],
    template: `
        <app-auth-layout>
            <router-outlet />
        </app-auth-layout>
    `
})
export class AuthWrapper {}
