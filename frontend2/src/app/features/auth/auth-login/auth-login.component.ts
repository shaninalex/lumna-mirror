import { Component, inject, signal } from "@angular/core";
import { FormsModule } from "@angular/forms";
import { actionSessionAuthenticateStart } from "@core/store";
import { Store } from "@ngrx/store";
import { form, required, email, FormField } from "@angular/forms/signals";

interface LoginFormPayload {
    email: string;
    password: string;
}

@Component({
    selector: "auth-login-feature",
    imports: [FormsModule, FormField],
    template: `
        <form (submit)="onSubmit()">
            <div class="mb-3">
                <label for="email" class="text-sm">Email address</label>
                <input
                    id="email"
                    class="form-control"
                    type="email"
                    placeholder="your@mail.org"
                    [formField]="loginForm.email"
                />
                @if (loginForm.email().dirty() && loginForm.email().errors()) {
                    @for (error of loginForm.email().errors(); track error) {
                        <div class="text-red-400 text-sm">
                            {{ error.message }}
                        </div>
                    }
                }
            </div>

            <div class="mb-3">
                <label for="password" class="text-sm">Password</label>
                <input
                    class="form-control"
                    id="password"
                    type="password"
                    placeholder="***"
                    [formField]="loginForm.password"
                />
                @if (
                    loginForm.password().dirty() &&
                    loginForm.password().errors()
                ) {
                    @for (error of loginForm.password().errors(); track error) {
                        <div class="text-red-400 text-sm">
                            {{ error.message }}
                        </div>
                    }
                }
            </div>
            <hr class="mb-3 border-slate-300" />
            <div>
                <button type="submit" class="btn">Login</button>
            </div>
        </form>
    `
})
export class AuthLoginFeature {
    private store = inject(Store);

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
