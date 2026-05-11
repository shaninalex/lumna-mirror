import { Component } from "@angular/core";
import { RouterOutlet } from "@angular/router";
import { WorkspaceHeader } from "./components";

@Component({
    selector: "app-workspace-wrapper",
    imports: [RouterOutlet, WorkspaceHeader],
    template: `
        <app-workspace-header />
        <router-outlet />
    `
})
export class WorkspaceWrapper {}
