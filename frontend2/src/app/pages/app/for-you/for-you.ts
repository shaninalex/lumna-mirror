import { Component } from "@angular/core";
import { RouterLink } from "@angular/router";

@Component({
    selector: "app-for-you",
    imports: [RouterLink],
    template: `
        <div class="p-4">
            <p>Latest assignments,</p>
            <div class="mb-4 flex items-start gap-2">
                <a
                    routerLink="/app/lumna-1"
                    class="border rounded p-4 hover:underline"
                >
                    Lumna ( 12 )
                </a>
                <a
                    routerLink="/app/new-frontend-1"
                    class="border rounded p-4 hover:underline"
                >
                    Another workpsace ( 2 )
                </a>
            </div>
            <div>
                <a routerLink="/workspaces" class="underline">View all</a>
            </div>
        </div>
    `
})
export class ForYou {}
