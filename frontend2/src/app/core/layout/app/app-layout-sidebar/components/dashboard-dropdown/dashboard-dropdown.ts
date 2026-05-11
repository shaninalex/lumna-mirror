import { Component } from "@angular/core";
import { CdkMenu, CdkMenuItem, CdkMenuTrigger } from "@angular/cdk/menu";
import { WorkspaceLabel, WorkspaceModel } from "@entities/workspace";

@Component({
    selector: "app-dashboard-dropdown",
    imports: [CdkMenuTrigger, CdkMenu, WorkspaceLabel], // CdkMenuItem
    template: `
        <button [cdkMenuTriggerFor]="dashboardDropdown">
            <app-workspace-label [workspace]="workspace" size="sm" />
        </button>
        <ng-template #dashboardDropdown>
            <div cdkMenu class="dropdown-base w-64 flex flex-col gap-2">
                <app-workspace-label [workspace]="workspace" size="md" />
                <div class="line-divider"></div>
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
