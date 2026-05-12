import { Component } from "@angular/core";
import { RouterLink } from "@angular/router";
import { WorkspaceLabel, WorkspaceModel } from "@entities/workspace";

@Component({
    selector: "app-workspace-list",
    template: `
        <div class="p-6">
            <h1 class="text-lg font-bold">Workspaces list</h1>
            <p class="text-sm text-slate-500 mb-4">
                This is the list of all your workspaces ( or workspaces you have
                permissions ).
            </p>
            <div class="flex flex-col gap-4">
                @for (workspace of workspaces; track $index) {
                    <a [routerLink]="['/app', workspace.slug]">
                        <app-workspace-label [workspace]="workspace" />
                    </a>
                }
            </div>
        </div>
    `,
    imports: [WorkspaceLabel, RouterLink]
})
export class WorkspaceList {
    workspaces: WorkspaceModel[] = [
        {
            id: 1,
            slug: "lumna-1",
            title: "Lumna",
            icon: "/img/project.svg"
        },
        {
            id: 2,
            slug: "new-frontend-1",
            title: "NewTestproject",
            icon: "/img/project.svg"
        }
    ];
}
