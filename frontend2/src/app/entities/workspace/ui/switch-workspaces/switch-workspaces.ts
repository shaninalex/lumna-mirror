import { Component } from "@angular/core";
import { RouterLink } from "@angular/router";
import { WorkspaceModel } from "@entities/workspace";

@Component({
    selector: "app-switch-workspaces",
    imports: [RouterLink],
    template: `
        <div>
            <div class="flex flex-col gap-2">
                @for (workspace of workspaces; track $index) {
                    <a
                        [routerLink]="['/app', workspace.slug]"
                        class="flex items-center gap-2"
                    >
                        <div>
                            <img
                                src="/img/project.svg"
                                alt=""
                                class="rounded w-5 h-5"
                            />
                        </div>
                        <div class="leading-none">
                            <div class="font-bold">
                                {{ workspace.title }}
                            </div>
                        </div>
                    </a>
                }
            </div>

            <a
                routerLink="/workspaces"
                class="text-slate-500 text-xs hover:underline"
            >
                view all
            </a>
        </div>
    `
})
export class SwitchWorkspaces {
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
