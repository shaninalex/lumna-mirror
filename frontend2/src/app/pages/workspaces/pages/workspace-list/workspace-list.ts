import { AsyncPipe } from "@angular/common";
import { Component, inject } from "@angular/core";
import { RouterLink } from "@angular/router";
import { selectWorkspaceList } from "@entities/workspace/model/workspace.selectors";
import { Store } from "@ngrx/store";
import { CardModule } from "primeng/card";
import { ButtonModule } from "primeng/button";

@Component({
    selector: "app-workspace-list",
    imports: [RouterLink, AsyncPipe, CardModule, ButtonModule],
    template: `
        <div class="workspace-layout">
            <div>
                <h1>Workspaces</h1>

                <p class="subtitle mb-4">
                    This is the list of all your workspaces (or workspaces you
                    have permissions for).
                </p>

                @if (workspaces$ | async; as workspaces) {
                    <div class="flex flex-col gap-2 mb-4">
                        @for (workspace of workspaces; track workspace.id) {
                            <p-card [header]="workspace.title">
                                <ng-template #footer>
                                    <p-button
                                        label="View"
                                        class="w-full"
                                        [outlined]="true"
                                        [routerLink]="['/app', workspace.id]"
                                    />
                                </ng-template>
                            </p-card>
                        }
                    </div>
                }

                <div class="flex gap-2">
                    <p-button routerLink="/app/workspaces/create">
                        Create workspace
                    </p-button>

                    <p-button
                        variant="outlined"
                        severity="secondary"
                        routerLink="/app/workspaces/archived"
                    >
                        View archived
                    </p-button>
                </div>
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
