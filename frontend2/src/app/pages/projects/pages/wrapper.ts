import { Component } from "@angular/core";
import { RouterOutlet } from "@angular/router";

@Component({
    selector: "app-project-wrapper",
    imports: [RouterOutlet],
    template: ` <router-outlet /> `
})
export class ProjectDetailWrapper {}
