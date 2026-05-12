import { Component } from "@angular/core";
import { RouterOutlet } from "@angular/router";

@Component({
    selector: "app-workspace-detail-wrapper",
    imports: [RouterOutlet],
    template: ` <router-outlet /> `
})
export class WorkspaceDetailWrapper {}
