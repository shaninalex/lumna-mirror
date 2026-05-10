import { Component } from "@angular/core";
import { RouterLink } from "@angular/router";

@Component({
    selector: "app-app-layout-sidebar",
    imports: [RouterLink],
    styleUrl: "./sidebar.css",
    template: ` <div class="sidebar">
        <div class="flex gap-4 flex-col">
            <a routerLink="/app/for-you">
                <img src="img/logo-h.svg" />
            </a>
            <nav class="flex gap-2">
                <a routerLink="/app" class="hover:underline">Spaces</a>
            </nav>
        </div>
    </div>`
})
export class AppLayoutSidebar {}
