import { AsyncPipe } from "@angular/common";
import { Component, inject, forwardRef } from "@angular/core";
import { RouterLink } from "@angular/router";
import { selectWorkspaceList } from "@entities/workspace/model/workspace.selectors";
import { Store } from "@ngrx/store";

@Component({
    selector: "app-for-you",
    imports: [RouterLink, AsyncPipe, forwardRef(() => NoWorkspaces)],
    template: `
        <div class="p-4">
            @if (workspaces$ | async; as workspaces) {
                @if (!workspaces.length) {
                    <app-no-workspaces />
                }
            } @else {
                <p>Latest assignments</p>
                <div>TODO: implement this</div>

                <div>
                    <a routerLink="/app/workspaces" class="underline">
                        View all workspaces
                    </a>
                </div>
            }
        </div>
    `
})
export class ForYou {
    private store = inject(Store);
    workspaces$ = this.store.select(selectWorkspaceList);
}

@Component({
    selector: "app-no-workspaces",
    imports: [RouterLink],
    template: `<div>
        You have no workspaces.
        <a routerLink="/app/workspaces/create" class="underline text-blue-500">
            Create one
        </a>
    </div> `
})
export class NoWorkspaces {}
