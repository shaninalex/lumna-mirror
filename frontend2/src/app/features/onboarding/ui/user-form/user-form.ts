import { Component, DestroyRef, inject, OnInit, signal } from "@angular/core";
import { takeUntilDestroyed } from "@angular/core/rxjs-interop";
import { email, form, FormField, required } from "@angular/forms/signals";
import { Store } from "@ngrx/store";
import {
    actionOnboardingCreateInvite,
    actionOnboardingCreateInviteFailed,
    actionOnboardingCreateInviteSuccess,
    UserOnboardingModel
} from "@features/onboarding/model";
import { Actions, ofType } from "@ngrx/effects";
import { tap } from "rxjs";
import { Error } from "@shared/models";

@Component({
    selector: "app-user-form-feature",
    imports: [FormField],
    template: `
        @if (success()) {
            <div class="text-green-600 mb-3">
                We sent invite on
                {{ onboardingForm.email().value() }}
            </div>
        } @else {
            <form (submit)="onSubmit($event)">
                <div class="mb-3">
                    <label class="form-label">Email address</label>
                    <input
                        type="email"
                        [formField]="onboardingForm.email"
                        class="form-control"
                        placeholder="name@example.com"
                    />
                </div>

                <div class="mb-3">
                    <label class="form-label">First name</label>
                    <input
                        type="text"
                        [formField]="onboardingForm.first_name"
                        class="form-control"
                        placeholder="First name"
                    />
                </div>

                <div class="mb-3">
                    <label class="form-label">Last name</label>
                    <input
                        type="text"
                        [formField]="onboardingForm.last_name"
                        class="form-control"
                        placeholder="Last name"
                    />
                </div>
                @for (error of errors(); track error) {
                    <div class="text-red-600 mb-3">{{ error.message }}</div>
                }
                <div>
                    <button class="btn btn-primary" type="submit">Next</button>
                </div>
            </form>
        }
    `
})
export class UserFormFeature {
    private store = inject(Store);
    private actions$ = inject(Actions);
    private destroyRef = inject(DestroyRef);

    errors = signal<Error[]>([]);
    success = signal(false);

    constructor() {
        this.actions$
            .pipe(
                ofType(actionOnboardingCreateInviteSuccess),
                tap(() => {
                    this.success.set(true);
                }),
                takeUntilDestroyed(this.destroyRef)
            )
            .subscribe();

        this.actions$
            .pipe(
                ofType(actionOnboardingCreateInviteFailed),
                tap(({ error }) => {
                    this.errors.set(error);
                }),
                takeUntilDestroyed(this.destroyRef)
            )
            .subscribe();
    }

    onboardingFormModel = signal<UserOnboardingModel>({
        email: "",
        first_name: "",
        last_name: ""
    });

    onboardingForm = form(this.onboardingFormModel, (schemaPath) => {
        required(schemaPath.first_name, { message: "First name is required" });
        required(schemaPath.last_name, { message: "Last name is required" });
        required(schemaPath.email, { message: "Email is required" });
        email(schemaPath.email, { message: "invalid email format" });
    });

    onSubmit(event: Event) {
        event.preventDefault();
        this.store.dispatch(
            actionOnboardingCreateInvite({
                payload: this.onboardingFormModel()
            })
        );
    }
}
