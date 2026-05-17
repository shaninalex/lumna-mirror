import { AsyncPipe } from "@angular/common";
import { Component, inject } from "@angular/core";
import { RouterLink } from "@angular/router";
import { selectWorkspaceList } from "@entities/workspace/model/workspace.selectors";
import { Store } from "@ngrx/store";

@Component({
    selector: "app-workspace-list",
    imports: [RouterLink, AsyncPipe],
    template: `
        <div class="p-4 h-screen flex flex-col justify-between">
            <div>
                <h1 class="text-lg font-bold">Workspaces list</h1>
                <p class="text-sm text-slate-500 mb-4">
                    This is the list of all your workspaces ( or workspaces you
                    have permissions ).
                </p>

                @if (workspaces$ | async; as workspaces) {
                    <div class="flex flex-col gap-4 mb-6">
                        @for (workspace of workspaces; track $index) {
                            <a
                                [routerLink]="['/app', workspace.slug]"
                                class="flex items-center gap-2"
                            >
                                <img
                                    src="/img/project.svg"
                                    alt=""
                                    class="w-10 h-10 rounded"
                                />
                                <span class="font-bold">
                                    {{ workspace.title }}
                                </span>
                            </a>
                        }
                    </div>
                }

                <a href="#" class="btn">Create</a>
            </div>
            <div class="hover:underline">Profile</div>
        </div>
    `
})
export class WorkspaceList {
    private store = inject(Store);
    workspaces$ = this.store.select(selectWorkspaceList);
}
