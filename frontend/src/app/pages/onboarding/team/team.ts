import { Component, inject, signal } from "@angular/core";
import { email, form, FormField, required } from "@angular/forms/signals";
import { FormsModule, ReactiveFormsModule } from "@angular/forms";
import { OnboardingApiService, TeamPageModel } from '@features/onboarding';

@Component({
    selector: "app-team-onboarding",
    imports: [FormsModule, ReactiveFormsModule, FormField],
    template: `
        <div class="card mb-4">
            <div class="card-body">
                <h1>Invite members</h1>

                <div class="d-flex flex-column gap-4">
                    <form (submit)="onSubmit()">
                        <div class="input-group">
                            <input
                                type="email"
                                [formField]="teamForm.email"
                                class="form-control"
                                placeholder="name@example.com"
                            />
                            <button class="input-group-text" type="button" (click)="addEmail()">+</button>
                        </div>
                        @if (
                            teamForm.email().dirty() &&
                            teamForm.email().errors()
                        ) {
                            @for (
                                error of teamForm.email().errors();
                                track error
                            ) {
                                <div class="text-danger small">
                                    {{ error.message }}
                                </div>
                            }
                        }

                        @if (teamEmails.length > 0) {
                            <ul class="list-group">
                                @for (email of teamEmails; track $index) {
                                    <li class="list-group-item">{{ email }}</li>
                                }
                            </ul>
                        }
                    </form>
                    <div>
                        <button class="btn btn-primary" type="submit">
                            Submit
                        </button>
                    </div>
                </div>
            </div>
        </div>

        <div class="text-center ">
            <button class="btn btn-outline-secondary" type="button">
                Skip
            </button>
        </div>
    `,
    providers: [OnboardingApiService],
})
export class TeamOnboardingPage {
    api = inject(OnboardingApiService);
    teamModel = signal({ email: "" });
    teamForm = form(this.teamModel, (schemaPath) => {
        required(schemaPath.email, { message: "Email is required" });
        email(schemaPath.email, { message: "Invalid email format" });
    });
    teamEmails: string[] = [];

    addEmail() {
        this.teamEmails.push(this.teamForm.email().value());
        this.teamForm.email().value.set("");
    }

    onSubmit(): void {
        const data: TeamPageModel = {
            emails: this.teamEmails,
        }
        this.api.team(data)
    }
}
