import { Component, inject, signal } from "@angular/core";
import { FormsModule } from "@angular/forms";
import { actionLoginFailed, actionSessionAuthenticateStart } from "@core/store";
import { Store } from "@ngrx/store";
import { email, form, FormField, required } from "@angular/forms/signals";
import { Actions, ofType } from "@ngrx/effects";
import { map } from "rxjs";
import { AsyncPipe } from "@angular/common";
import { RouterLink } from "@angular/router";
import { InputTextModule } from "primeng/inputtext";
import { FloatLabelModule } from "primeng/floatlabel";
import { ButtonModule } from "primeng/button";
import { MessageModule } from "primeng/message";
import { CardModule } from "primeng/card";

interface LoginFormPayload {
    email: string;
    password: string;
}

@Component({
    selector: "auth-login-feature",
    imports: [
        FormsModule,
        FormField,
        AsyncPipe,
        RouterLink,
        InputTextModule,
        FloatLabelModule,
        ButtonModule,
        MessageModule,
        CardModule
    ],
    templateUrl: "./auth-login.component.html"
})
export class AuthLoginFeature {
    private store = inject(Store);
    private actions$ = inject(Actions);
    errors$ = this.actions$.pipe(
        ofType(actionLoginFailed),
        map((action) => action.errors)
    );

    loginFormModel = signal<LoginFormPayload>({
        email: "",
        password: ""
    });

    loginForm = form(this.loginFormModel, (schemaPath) => {
        required(schemaPath.email, { message: "Email is required" });
        required(schemaPath.password, { message: "Password is required" });
        email(schemaPath.email, { message: "invalid email format" });
    });

    onSubmit(): void {
        if (
            !this.loginForm.email().errors().length &&
            !this.loginForm.password().errors().length
        ) {
            this.store.dispatch(
                actionSessionAuthenticateStart({
                    email: this.loginForm.email().value(),
                    password: this.loginForm.password().value()
                })
            );
        }
    }
}
