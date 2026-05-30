import { Component, inject } from "@angular/core";
import { MenuModule } from "primeng/menu";
import { MenuItem, MessageService } from "primeng/api";
import { PanelMenuModule } from "primeng/panelmenu";

@Component({
    selector: "app-app-layout-sidebar",
    imports: [MenuModule, PanelMenuModule],
    providers: [MessageService],
    styleUrl: "./app-layout-sidebar.css",
    template: `
        <p-panelmenu
            [model]="projectSettings"
            class="sidebar-menu"
        ></p-panelmenu>
        <p-menu [model]="items" class="sidebar-menu" />
    `
})
export class AppLayoutSidebar {
    private messageService = inject(MessageService);
    projectSettings: MenuItem[];
    items: MenuItem[] | undefined;

    ngOnInit() {
        this.projectSettings = [
            {
                label: "Project settings",
                icon: "pi pi-cog",
                items: [
                    { label: "Settings", icon: "pi pi-cog" },
                    { label: "Members", icon: "pi pi-cog" },
                    { label: "Permissions", icon: "pi pi-cog" },
                    { label: "Integrations", icon: "pi pi-cog" },
                    { label: "Webhooks", icon: "pi pi-cog" }
                ]
            }
        ];
        this.items = [
            // {
            //     separator: true
            // },
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
