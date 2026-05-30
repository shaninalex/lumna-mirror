import { Component, inject } from "@angular/core";
import { selectWorkspaceList } from "@entities/workspace/model/workspace.selectors";
import { Store } from "@ngrx/store";
import { AsyncPipe } from "@angular/common";
import { RouterLink } from "@angular/router";

@Component({
    selector: "app-switch-workspaces",
    imports: [AsyncPipe, RouterLink],
    template: `
        <button aria-label="Switch workspaces dropdown menu">
            Switch workspaces
        </button>
        <ng-template>
            @if (workspaces$ | async; as workspaces) {
                @for (workspace of workspaces; track $index) {
                    <a [routerLink]="['/app', workspace.slug]">
                        {{ workspace.title }}
                    </a>
                }
                <hr class="divider" />
                <a routerLink="/app/workspaces"> view all </a>
            }
        </ng-template>
    `
})
export class SwitchWorkspaces {
    private store = inject(Store);
    workspaces$ = this.store.select(selectWorkspaceList);
}
