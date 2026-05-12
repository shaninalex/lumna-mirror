import { Component } from "@angular/core";
import { RouterLink, RouterLinkActive } from "@angular/router";

@Component({
    selector: "app-project-header",
    imports: [RouterLink, RouterLinkActive],
    template: `
        <div class="p-4 border-b  border-slate-200">
            <div class="mb-3">
                <div class="flex gap-2 items-center">
                    <img
                        src="/img/project.svg"
                        class="rounded"
                        style="width: 25px"
                    />
                    <div class="font-bold text-lg">project name</div>
                </div>
            </div>
            <nav class="flex gap-4">
                <div>
                    <a
                        routerLink="/app/lumna-1/project/lumna-new-frontend-13/summary"
                        [routerLinkActive]="['underline', 'text-blue-500']"
                        class="hover:underline"
                    >
                        Summary
                    </a>
                </div>
                <div>
                    <a
                        routerLink="/app/lumna-1/project/lumna-new-frontend-13/backlog"
                        [routerLinkActive]="['underline', 'text-blue-500']"
                        class="hover:underline"
                    >
                        Backlog
                    </a>
                </div>

                <div>
                    <a
                        routerLink="/app/lumna-1/project/lumna-new-frontend-13/board"
                        [routerLinkActive]="['underline', 'text-blue-500']"
                        class="hover:underline"
                    >
                        Board
                    </a>
                </div>
            </nav>
        </div>
    `
})
export class ProjectHeader {}
