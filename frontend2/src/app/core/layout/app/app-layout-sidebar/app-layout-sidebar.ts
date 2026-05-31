import { Component } from "@angular/core";
import { MenuModule } from "primeng/menu";
import { MessageService } from "primeng/api";
import { PanelMenuModule } from "primeng/panelmenu";
import { PopoverModule } from "primeng/popover";
import { ButtonModule } from "primeng/button";
import { SwitchWorkspaces, WorkspaceSettings } from "@entities/workspace";
import { DividerModule } from "primeng/divider";
import { ProjectList } from "@entities/project";

@Component({
    selector: "app-app-layout-sidebar",
    imports: [
        MenuModule,
        PanelMenuModule,
        PopoverModule,
        ButtonModule,
        SwitchWorkspaces,
        DividerModule,
        WorkspaceSettings,
        ProjectList
    ],
    providers: [MessageService],
    template: `
        <div class="flex flex-col h-full">
            <app-project-list />

            <div class="grow"></div>

            <p-divider />
            <div class="px-4 pb-4 flex flex-col gap-1">
                <app-workspace-settings />
                <app-switch-workspaces />
            </div>
        </div>
    `
})
export class AppLayoutSidebar {}
