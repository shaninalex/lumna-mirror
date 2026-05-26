import { Component } from "@angular/core";
import { WorkspaceCreateFeature } from "@features";

@Component({
    selector: "app-workspace-create-page",
    imports: [WorkspaceCreateFeature],
    template: `<app-workspace-create-feature />`
})
export class WorkspaceCreatePage {}
