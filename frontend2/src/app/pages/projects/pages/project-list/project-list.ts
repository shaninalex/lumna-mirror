import { Component } from "@angular/core";
import { RouterLink } from "@angular/router";

@Component({
    selector: "app-project-list",
    imports: [RouterLink],
    template: `<div class="p-4">
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
            <a
                routerLink="/app/lumna-1/project/exelunatic-33"
                class="block hover:underline"
            >
                Project C
            </a>
        </nav>
        <a
            routerLink="/app/lumna-1/projects"
            class="text-slate-500 text-xs hover:underline"
        >
            Create
        </a>
    </div>`
})
export class ProjectListPage {}
