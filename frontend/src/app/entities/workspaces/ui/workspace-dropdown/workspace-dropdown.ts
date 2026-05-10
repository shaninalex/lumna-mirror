import { Component, Inject, inject, OnInit } from "@angular/core";
import { CdkMenu, CdkMenuItem, CdkMenuTrigger } from "@angular/cdk/menu";
import { Store } from "@ngrx/store";
import { selectWorkspaceList } from "@entities/workspaces";
import { AsyncPipe } from "@angular/common";

@Component({
    selector: "app-workspace-dropdown",
    imports: [CdkMenuTrigger, CdkMenu, CdkMenuItem, AsyncPipe],
    template: `
        @if (workspaces$ | async; as workspaces) {
            <div class="dropdown">
                <button
                    class="btn btn-secondary dropdown-toggle"
                    type="button"
                    [cdkMenuTriggerFor]="menu"
                    aria-expanded="false"
                >
                    Dropdown
                </button>
                <ng-template #menu>
                    <ul class="dropdown-menu d-block" cdkMenu>
                        @for (workspace of workspaces; track workspace.id) {
                            <button
                                class="dropdown-item"
                                cdkMenuItem
                                type="button"
                            >
                                {{ workspace.title }}
                            </button>
                        }
                    </ul>
                </ng-template>
            </div>
        } @else {
            loading workspaces...
        }
    `
})
export class WorkspaceDropdown {
    store = inject(Store);
    workspaces$ = this.store.select(selectWorkspaceList);
}
