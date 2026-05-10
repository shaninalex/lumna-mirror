import { Component } from "@angular/core";
import { RouterLink } from "@angular/router";

@Component({
    selector: "app-workspaces-list",
    imports: [RouterLink],
    template: `
        <ul>
            <li>
                <a routerLink="/app/workspace/1" class="underline"
                    >Workspace A</a
                >
            </li>
            <li>
                <a routerLink="/app/workspace/2" class="underline"
                    >Workspace B</a
                >
            </li>
            <li>
                <a routerLink="/app/workspace/3" class="underline"
                    >Workspace C</a
                >
            </li>
            <li>
                <a routerLink="/app/workspace/4" class="underline"
                    >Workspace D</a
                >
            </li>
        </ul>
    `
})
export class WorkspacesList {}
