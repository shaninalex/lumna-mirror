import { Component, OnInit } from "@angular/core";
import { MenuModule } from "primeng/menu";
import { MenuItem, MessageService } from "primeng/api";
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
    styleUrl: "./app-layout-sidebar.css",
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
export class AppLayoutSidebar implements OnInit {
    items: MenuItem[] | undefined;

    ngOnInit() {
        this.items = [
            {
                label: "Projects",
                items: [
                    {
                        label: "Project name b",
                        icon: "pi pi-file",
                        routerLink: "/app/lumna-1/project/lumna-new-frontend-13"
                    },
                    {
                        label: "Project name a",
                        icon: "pi pi-file",
                        routerLink: "/app/lumna-1/project/sdondford-22"
                    }
                ]
            }
        ];
    }
}

/*
<div class="sidebar">
    <div class="">
        <app-dashboard-dropdown />
        <div>
            <a routerLink="/app/lumna-1/project/lumna-new-frontend-13">
                <span>Project A</span>
            </a>
            <a routerLink="/app/lumna-1/project/sdondford-22">
                <span>Project B</span>
            </a>
        </div>
    </div>
    <app-switch-workspaces />
</div>
*/
