import { Component } from "@angular/core";
import { RouterOutlet } from "@angular/router";
import { ProjectHeader } from "./components";

@Component({
    selector: "app-project-wrapper",
    imports: [RouterOutlet, ProjectHeader],
    template: `
        <app-project-header />
        <router-outlet />
    `
})
export class ProjectWrapper {}
