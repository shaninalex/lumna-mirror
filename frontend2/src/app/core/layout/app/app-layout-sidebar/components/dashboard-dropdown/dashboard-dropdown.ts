import { Component } from "@angular/core";
import { CdkMenu, CdkMenuTrigger } from "@angular/cdk/menu";
import {
    WorkspaceLabel,
    WorkspaceModel,
    SwitchWorkspaces
} from "@entities/workspace";

@Component({
    selector: "app-dashboard-dropdown",
    imports: [CdkMenuTrigger, CdkMenu, WorkspaceLabel, SwitchWorkspaces], // CdkMenuItem
    template: `
        <button [cdkMenuTriggerFor]="dashboardDropdown">
            <app-workspace-label [workspace]="workspace" size="sm" />
        </button>
        <ng-template #dashboardDropdown>
            <div cdkMenu class="dropdown-base w-64 flex flex-col gap-2">
                <app-workspace-label [workspace]="workspace" size="md" />
                <div class="line-divider"></div>
                <ul>
                    <li>Settings</li>
                    <li>Members</li>
                </ul>
                <div class="line-divider"></div>
                <app-switch-workspaces />
            </div>
        </ng-template>
    `
})
export class DashboardDropdown {
    wspTitle = "Lumna";
    workspace: WorkspaceModel = {
        id: 123,
        slug: "lumna-123",
        title: "Lumna",
        icon: "/img/project.svg"
    };

    textChanged(): void {
        console.log("on click");
    }
}
