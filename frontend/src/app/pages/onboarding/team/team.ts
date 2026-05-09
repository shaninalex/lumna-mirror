import { Component, signal } from "@angular/core";
import { email, form, FormField, required } from '@angular/forms/signals';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

@Component({
    selector: "app-team-onboarding",
    imports: [
        FormsModule,
        ReactiveFormsModule,
        FormField
    ],
    template: `
        <h1>Invite members</h1>

        <div class="d-flex flex-column gap-4">
            <form (submit)="onSubmit()">
                <div class="mb-3">
                    <label class="form-label">Email</label>
                    <input type="email"
                           [formField]="teamForm.email"
                           class="form-control"
                           placeholder="name@example.com">
                    @if (teamForm.email().dirty() && teamForm.email().errors()) {
                        @for (error of teamForm.email().errors(); track error) {
                            <div class="text-danger small">{{ error.message }}</div>
                        }
                    }
                </div>
                <div>
                    <button class="btn btn-primary" type="submit">Add</button>
                </div>
            </form>

            @if (teamEmails.length > 0)  {
                <ul class="list-group">
                    @for (email of teamEmails; track $index) {
                        <li class="list-group-item">{{ email }}</li>
                    }
                </ul>
            }
        </div>
    `
})
export class TeamOnboardingPage {
    teamModel = signal({ email: '' });
    teamForm = form(this.teamModel, schemaPath => {
        required(schemaPath.email, { message: "Email is required"});
        email(schemaPath.email, { message: "Invalid email format"});
    });
    teamEmails: string[] = []

    onSubmit() {
        this.teamEmails.push(this.teamForm.email().value())
        this.teamForm.email().value.set('')
    }
}
