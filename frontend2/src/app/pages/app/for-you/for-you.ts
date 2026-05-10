import { Component } from "@angular/core";
import { RouterLink } from "@angular/router";

@Component({
    selector: "app-for-you",
    imports: [RouterLink],
    template: `
        <div class="mb-4 flex items-start gap-4">
            <a
                routerLink="/app/workspace/1"
                class="border rounded p-4 hover:underline"
            >
                Workspace a
            </a>
            <a
                routerLink="/app/workspace/2"
                class="border rounded p-4 hover:underline   "
            >
                Workspace b
            </a>
        </div>
        <div>
            <a routerLink="/app" class="underline">View all workspaces</a>
        </div>
    `
})
export class ForYou {}
