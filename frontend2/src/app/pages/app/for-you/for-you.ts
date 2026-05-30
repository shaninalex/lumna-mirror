import { AsyncPipe } from "@angular/common";
import { Component, inject } from "@angular/core";
import { RouterLink } from "@angular/router";
import { selectWorkspaceList } from "@entities/workspace/model/workspace.selectors";
import { Store } from "@ngrx/store";
import { CardModule } from "primeng/card";

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

@Component({
    selector: "app-for-you",
    imports: [RouterLink, AsyncPipe, NoWorkspaces, CardModule],
    template: `
        <div class="p-16 flex justify-center w-full">
            @if (workspaces$ | async; as workspaces) {
                @if (!workspaces.length) {
                    <app-no-workspaces />
                } @else {
                    <p-card header="Today's tasks">
                        @for (workspace of workspaces; track workspace.id) {
                            <a
                                [routerLink]="['/app', workspace.slug]"
                                class="flex flex-col gap-2 p-4 cursor-pointer hover:bg-emphasis transition-all no-underline"
                            >
                                <span class="text-color font-medium leading-6">
                                    {{ workspace.title }}
                                </span>
                                <span
                                    class="text-muted-color text-sm leading-5 line-clamp-1"
                                >
                                    Tasks in progress: 12, total: 29
                                </span>
                            </a>
                        }
                    </p-card>
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
