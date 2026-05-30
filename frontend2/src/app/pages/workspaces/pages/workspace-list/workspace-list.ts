import { AsyncPipe } from "@angular/common";
import { Component, inject } from "@angular/core";
import { RouterLink } from "@angular/router";
import { selectWorkspaceList } from "@entities/workspace/model/workspace.selectors";
import { Store } from "@ngrx/store";

@Component({
    selector: "app-workspace-list",
    imports: [RouterLink, AsyncPipe],
    template: `
        <div class="workspace-layout">
            <div>
                <h1>Workspaces</h1>

                <p class="subtitle">
                    This is the list of all your workspaces (or workspaces you
                    have permissions for).
                </p>

                @if (workspaces$ | async; as workspaces) {
                    <div>
                        @for (workspace of workspaces; track workspace.id) {
                            <a [routerLink]="['/app', workspace.slug]">
                                <img src="/img/project.svg" alt="" />
                                <span>
                                    {{ workspace.title }}
                                </span>
                            </a>
                        }
                    </div>
                }

                <button routerLink="/app/workspaces/create">
                    Create workspace
                </button>
            </div>

            <div>
                <button>Profile</button>
            </div>
        </div>
    `,
    styleUrl: "./workspace-list.css"
})
export class WorkspaceList {
    private store = inject(Store);
    workspaces$ = this.store.select(selectWorkspaceList);
}
