import { Component, inject, signal } from "@angular/core";
import { email, form, FormField, required, schema } from "@angular/forms/signals";
import { TeamOnboardingPage } from "@pages/onboarding/team/team";
import { OnboardingApiService, WorkspacePageModel } from '@features/onboarding';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { Router } from '@angular/router';
import { KanbanApi } from '@features/kanban-board';

@Component({
    selector: "app-workspace-onboarding",
    imports: [
        FormsModule,
        ReactiveFormsModule,
        FormField
    ],
    templateUrl: "./workspace.html",
    providers: [OnboardingApiService],
})
export class WorkspaceOnboardingPage {
    api = inject(OnboardingApiService)
    router = inject(Router)

    workspaceFormModel = signal<WorkspacePageModel>({
        title: "",
        email: ""
    });
    workspaceForm = form(this.workspaceFormModel, (schemaPath) => {
        required(schemaPath.email);
        email(schemaPath.email);
        required(schemaPath.title);
    });

    onSubmit(): void {
        const data = this.workspaceForm().value()
        this.api.workspace(data).subscribe(data => {
            console.log(data);
            this.router.navigateByUrl("/onboarding/team");
        })
    }
}
