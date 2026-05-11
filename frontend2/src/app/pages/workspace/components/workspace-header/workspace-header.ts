import { Component } from "@angular/core";
import { RouterLink, RouterLinkActive } from "@angular/router";

@Component({
    selector: "app-workspace-header",
    imports: [RouterLink, RouterLinkActive],
    template: `
        <div class="p-4 border-b  border-slate-200">
            <div class="mb-3">
                <div class="text-slate-600 text-sm">Workspaces</div>
                <div class="flex gap-2 items-center">
                    <img
                        src="/img/project.svg"
                        class="rounded"
                        style="width: 25px"
                    />
                    <div class="font-bold text-lg">Workspace name</div>
                </div>
            </div>
            <nav class="flex gap-4">
                <a
                    routerLink="summary"
                    [routerLinkActive]="['underline', 'text-blue-500']"
                    class="hover:underline"
                >
                    Summary
                </a>
                <a
                    routerLink="backlog"
                    [routerLinkActive]="['underline', 'text-blue-500']"
                    class="hover:underline"
                >
                    Backlog
                </a>
                <a
                    routerLink="board"
                    [routerLinkActive]="['underline', 'text-blue-500']"
                    class="hover:underline"
                >
                    Board
                </a>
            </nav>
        </div>
    `
})
export class WorkspaceHeader {}
