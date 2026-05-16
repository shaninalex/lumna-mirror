import { Component } from "@angular/core";
import { RouterLink } from "@angular/router";
import { DashboardDropdown, ToggleSidebar } from "./components";

@Component({
    selector: "app-app-layout-sidebar",
    imports: [RouterLink, DashboardDropdown, ToggleSidebar],
    styleUrl: "./sidebar.css",
    template: ` <div class="sidebar">
        <div class="flex gap-2 flex-col">
            <div class="flex gap-2 items-center justify-between">
                <app-dashboard-dropdown />
                <app-toggle-sidebar />
            </div>

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
export class AppLayoutSidebar {}
