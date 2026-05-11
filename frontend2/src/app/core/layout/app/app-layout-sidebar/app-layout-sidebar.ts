import { Component } from "@angular/core";
import { RouterLink } from "@angular/router";
import { DashboardDropdown } from "./components";

@Component({
    selector: "app-app-layout-sidebar",
    imports: [RouterLink, DashboardDropdown],
    styleUrl: "./sidebar.css",
    template: ` <div class="sidebar">
        <div class="flex gap-2 flex-col">
            <app-dashboard-dropdown />
            <nav class="flex gap-2 flex-col">
                <div class="text-slate-500 text-sm">Projects</div>
                <a routerLink="/app/project/1" class="block hover:underline">
                    Project A
                </a>
                <a routerLink="/app/project/2" class="block hover:underline">
                    Project B
                </a>
                <a routerLink="/app/project/3" class="block hover:underline">
                    Project C
                </a>
                <a
                    routerLink="/app/projects"
                    class="text-slate-500 text-xs hover:underline"
                >
                    view all
                </a>

                <hr class="border-slate-200" />
            </nav>
        </div>
    </div>`
})
export class AppLayoutSidebar {}
