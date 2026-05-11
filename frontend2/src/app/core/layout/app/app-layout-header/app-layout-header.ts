import { Component } from "@angular/core";

@Component({
    selector: "app-app-layout-header",
    imports: [],
    styleUrl: "./header.css",
    template: `
        <header class="header">
            <a routerLink="/app/for-you" class="mb-4">
                <img src="img/logo-h.svg" />
            </a>
        </header>
    `
})
export class AppLayoutHeader {}
