import { Component } from "@angular/core";
import { CdkMenu, CdkMenuItem, CdkMenuTrigger } from "@angular/cdk/menu";
import { SwitchWorkspaces, WorkspaceModel } from "@entities/workspace";

@Component({
    selector: "app-dashboard-dropdown",
    imports: [CdkMenuTrigger, CdkMenu, CdkMenuItem, SwitchWorkspaces],
    template: `
        <button
            [cdkMenuTriggerFor]="dashboardDropdown"
            class="inline-block cursor-pointer p-1 rounded hover:bg-slate-300 cursor-pointer;"
        >
            <!--            <app-workspace-label [workspace]="workspace" size="sm" />-->
        </button>
        <ng-template #dashboardDropdown>
            <div cdkMenu class="dropdown-base w-64 flex flex-col gap-2">
                <!--                <app-workspace-label [workspace]="workspace" size="md" />-->
                <div class="line-divider"></div>
                <ul>
                    <li cdkMenuItem>Settings</li>
                    <li cdkMenuItem>Members</li>
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
}
