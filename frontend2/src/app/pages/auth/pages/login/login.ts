import { Component } from "@angular/core";
import { AuthLoginFeature } from "@features/auth";

@Component({
    selector: "app-login",
    imports: [AuthLoginFeature],
    template: `<auth-login-feature />
    `
})
export class LoginPage {}
