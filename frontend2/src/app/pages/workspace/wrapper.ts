import { Component } from "@angular/core";
import { RouterLink, RouterLinkActive, RouterOutlet } from "@angular/router";

@Component({
    selector: "app-workspace-wrapper",
    imports: [RouterOutlet, RouterLink, RouterLinkActive],
    template: `
        <nav class="flex gap-2 border-b mb-3 pb-2 border-slate-200">
            <a
                routerLink="summary"
                [routerLinkActive]="['underline']"
                class="hover:underline"
            >
                Summary
            </a>
            <a
                routerLink="backlog"
                [routerLinkActive]="['underline']"
                class="hover:underline"
            >
                Backlog
            </a>
            <a
                routerLink="board"
                [routerLinkActive]="['underline']"
                class="hover:underline"
            >
                Board
            </a>
        </nav>
        <router-outlet />
    `
})
export class WorkspaceWrapper {}
