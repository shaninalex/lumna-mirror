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
                <a
                    routerLink="/app/lumna-1/project/exelunatic-33"
                    class="block hover:underline"
                >
                    Project C
                </a>
                <a
                    routerLink="/app/lumna-1/projects"
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
