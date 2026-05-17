import { Component, inject, signal } from "@angular/core";
import { email, form, FormField, required } from "@angular/forms/signals";
import { OnboardingApiService } from "@features/onboarding";
import { UserOnboardingModel } from "@features/onboarding/model";

@Component({
    selector: "app-user-form-feature",
    imports: [FormField],
    providers: [OnboardingApiService],
    template: `
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

            <div>
                <button class="btn btn-primary" type="submit">Next</button>
            </div>
        </form>
    `
})
export class UserFormFeature {
    private api = inject(OnboardingApiService);
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
        this.api.user(this.onboardingFormModel()).subscribe();
    }
}
