import { Component } from "@angular/core";
import { WorkspaceArchivedFeature } from "@features/workspace";

@Component({
    selector: "app-workspace-archived-page",
    imports: [WorkspaceArchivedFeature],
    template: `<app-workspace-archived-feature />`
})
export class WorkspaceArchivedPage {}
