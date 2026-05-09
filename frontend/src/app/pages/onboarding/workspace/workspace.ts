import { Component, signal } from "@angular/core";
import { email, form, required, schema } from "@angular/forms/signals";
import { TeamOnboardingPage } from "@pages/onboarding/team/team";

@Component({
    selector: "app-workspace-onboarding",
    imports: [],
    templateUrl: "./workspace.html"
})
export class WorkspaceOnboardingPage {
    workspaceFormModel = signal({
        title: "",
        email: ""
    });
    workspaceForm = form(this.workspaceFormModel, (schemaPath) => {
        required(schemaPath.email);
        email(schemaPath.email);
        required(schemaPath.title);
    });
}
