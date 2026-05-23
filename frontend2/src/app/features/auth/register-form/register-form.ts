import { Component, inject, Input, OnInit, signal, input } from "@angular/core";
import { AsyncPipe } from "@angular/common";
import { FormsModule } from "@angular/forms";
import { form, required, validate, FormField } from "@angular/forms/signals";
import { Actions, ofType } from "@ngrx/effects";
import { Store } from "@ngrx/store";
import { map } from "rxjs";

import { actionLoginFailed } from "@core";
import {
    AcceptInviteUserRegistrationForm,
    actionUserInviteRegister
} from "@entities/user";
import { OnboardingContinue } from "@features/onboarding";

@Component({
    selector: "app-register-form-feature",
    imports: [FormsModule, FormField, AsyncPipe],
    template: `
        <form (submit)="onSubmit()">
            <div class="mb-3">
                <label for="first_name" class="text-sm"> First name </label>

                <input
                    id="first_name"
                    class="form-control"
                    type="text"
                    placeholder="John"
                    [formField]="registerForm.first_name"
                />

                @if (
                    registerForm.first_name().dirty() &&
                    registerForm.first_name().errors()
                ) {
                    @for (
                        error of registerForm.first_name().errors();
                        track error
                    ) {
                        <div class="text-red-400 text-sm">
                            {{ error.message }}
                        </div>
                    }
                }
            </div>

            <div class="mb-3">
                <label for="last_name" class="text-sm"> Last name </label>

                <input
                    id="last_name"
                    class="form-control"
                    type="text"
                    placeholder="Doe"
                    [formField]="registerForm.last_name"
                />

                @if (
                    registerForm.last_name().dirty() &&
                    registerForm.last_name().errors()
                ) {
                    @for (
                        error of registerForm.last_name().errors();
                        track error
                    ) {
                        <div class="text-red-400 text-sm">
                            {{ error.message }}
                        </div>
                    }
                }
            </div>

            <div class="mb-3">
                <label for="password" class="text-sm"> Password </label>

                <input
                    id="password"
                    class="form-control"
                    type="password"
                    placeholder="***"
                    [formField]="registerForm.password"
                />

                @if (
                    registerForm.password().dirty() &&
                    registerForm.password().errors()
                ) {
                    @for (
                        error of registerForm.password().errors();
                        track error
                    ) {
                        <div class="text-red-400 text-sm">
                            {{ error.message }}
                        </div>
                    }
                }
            </div>

            <div class="mb-3">
                <label for="password_confirm" class="text-sm">
                    Confirm password
                </label>

                <input
                    id="password_confirm"
                    class="form-control"
                    type="password"
                    placeholder="***"
                    [formField]="registerForm.password_confirm"
                />

                @if (
                    registerForm.password_confirm().dirty() &&
                    registerForm.password_confirm().errors()
                ) {
                    @for (
                        error of registerForm.password_confirm().errors();
                        track error
                    ) {
                        <div class="text-red-400 text-sm">
                            {{ error.message }}
                        </div>
                    }
                }
            </div>

            @if (registerForm().errors(); as errors) {
                @for (error of errors; track error) {
                    <div class="text-red-500 text-sm mb-2">
                        {{ error.message }}
                    </div>
                }
            }

            @if (errors$ | async; as errors) {
                @for (error of errors; track error) {
                    <div class="text-red-600 mb-3">
                        {{ error.message }}
                    </div>
                }
            }

            <div>
                <button type="submit" class="btn">Register</button>
            </div>
        </form>
    `
})
export class RegisterForm implements OnInit {
    @Input({ required: true }) invitationToken!: string;
    invitation = input<OnboardingContinue>();

    private store = inject(Store);
    private actions$ = inject(Actions);

    errors$ = this.actions$.pipe(
        ofType(actionLoginFailed),
        map((action) => action.errors)
    );

    registerFormModel = signal<
        AcceptInviteUserRegistrationForm & {
            password_confirm: string;
        }
    >({
        password: "",
        password_confirm: "",
        first_name: "",
        last_name: "",
        invitation_token: ""
    });

    registerForm = form(this.registerFormModel, (schemaPath) => {
        required(schemaPath.password, {
            message: "Password is required"
        });

        required(schemaPath.password_confirm, {
            message: "Password confirmation is required"
        });

        required(schemaPath.first_name, {
            message: "First name is required"
        });

        required(schemaPath.last_name, {
            message: "Last name is required"
        });

        required(schemaPath.invitation_token, {
            message: "Invitation token is required"
        });

        validate(schemaPath, ({ value }) => {
            if (value().password !== value().password_confirm) {
                return {
                    kind: "password_mismatch",
                    message: "Confirmation password should match password"
                };
            }

            return null;
        });
    });

    ngOnInit() {
        this.registerForm.invitation_token().value.set(this.invitationToken);
        const v = this.invitation();
        if (v) {
            this.registerForm.first_name().value.set(v.meta.first_name);
            this.registerForm.last_name().value.set(v.meta.last_name);
        }
    }

    onSubmit(): void {
        if (this.registerForm().errors().length) {
            return;
        }

        this.store.dispatch(
            actionUserInviteRegister({ data: this.registerFormModel() })
        );
    }
}
