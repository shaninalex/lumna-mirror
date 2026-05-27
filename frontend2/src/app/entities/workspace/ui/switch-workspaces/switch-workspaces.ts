import { Component, inject } from "@angular/core";
import { RouterLink } from "@angular/router";
import { CdkMenuItem } from "@angular/cdk/menu";
import { selectWorkspaceList } from "@entities/workspace/model/workspace.selectors";
import { Store } from "@ngrx/store";
import { AsyncPipe } from "@angular/common";

@Component({
    selector: "app-switch-workspaces",
    imports: [RouterLink, CdkMenuItem, AsyncPipe],
    template: `
        @if (workspaces$ | async; as workspaces) {
            <div>
                <div class="flex flex-col gap-2">
                    @for (workspace of workspaces; track $index) {
                        <a
                            [routerLink]="['/app', workspace.slug]"
                            class="flex items-center gap-2"
                            cdkMenuItem
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
                    cdkMenuItem
                    routerLink="/app/workspaces"
                    class="text-slate-500 text-xs hover:underline"
                >
                    view all
                </a>
            </div>
        }
    `
})
export class SwitchWorkspaces {
    private store = inject(Store);
    workspaces$ = this.store.select(selectWorkspaceList);
}
