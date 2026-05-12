import { Component } from "@angular/core";
import { RouterOutlet } from "@angular/router";
import { AppLayout } from "@core/layout";

@Component({
    selector: "application-wrapper",
    imports: [RouterOutlet, AppLayout],
    template: `
        <app-layout>
            <router-outlet />
        </app-layout>
    `
})
export class ApplicationWrapper {}
