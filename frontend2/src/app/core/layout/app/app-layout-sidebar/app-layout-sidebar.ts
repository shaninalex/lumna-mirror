import { Component, inject, OnInit } from "@angular/core";
import {
    NavigationEnd,
    Router,
    RouterLink,
    Event,
    NavigationStart
} from "@angular/router";
import { DashboardDropdown } from "./components";

@Component({
    selector: "app-app-layout-sidebar",
    imports: [RouterLink, DashboardDropdown],
    styleUrl: "./sidebar.css",
    template: ` <div class="sidebar">
        <div class="flex gap-2 flex-col">
            <app-dashboard-dropdown />

            <div class="text-slate-500 text-sm">Projects</div>
            <nav class="flex gap-2 flex-col">
                <a
                    routerLink="/app/lumna-1/project/lumna-new-frontend-13"
                    class="block hover:underline"
                >
                    Project A
                </a>
                <a
                    routerLink="/app/lumna-1/project/sdondford-22"
                    class="block hover:underline"
                >
                    Project B
                </a>
            </nav>
            <a
                routerLink="/app/lumna-1/projects"
                class="text-slate-500 text-xs hover:underline"
            >
                view all
            </a>

            <hr class="border-slate-200" />

            <nav class="flex gap-2 flex-col">
                <a
                    routerLink="/app/lumna-1/project/lumna-new-frontend-13"
                    class="block hover:underline"
                >
                    Settings
                </a>
            </nav>
        </div>

        <nav class="flex gap-2 flex-col">
            <div>
                <button class="block hover:underline">Profile</button>
            </div>
        </nav>
    </div>`
})
export class AppLayoutSidebar {
    private readonly router = inject(Router);

    constructor() {
        this.router.events.subscribe((event) => {
            if (event instanceof NavigationEnd) {
                // Navigation completed
                console.log("Navigation completed:", event);
            }
        });
    }
}
