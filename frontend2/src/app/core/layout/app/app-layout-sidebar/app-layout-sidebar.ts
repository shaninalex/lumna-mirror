import { Component, inject, OnInit } from "@angular/core";
import { MenuModule } from "primeng/menu";
import { MenuItem, MessageService } from "primeng/api";
import { PanelMenuModule } from "primeng/panelmenu";
import { PopoverModule } from "primeng/popover";
import { ButtonModule } from "primeng/button";
import { SwitchWorkspaces } from "@entities/workspace";

@Component({
    selector: "app-app-layout-sidebar",
    imports: [
        MenuModule,
        PanelMenuModule,
        PopoverModule,
        ButtonModule,
        SwitchWorkspaces
    ],
    providers: [MessageService],
    styleUrl: "./app-layout-sidebar.css",
    template: `
        <div class="flex flex-col h-full">
            <div class="p-4 pb-0 flex justify-start">
                <p-button
                    (click)="menu.toggle($event)"
                    label="Project settings"
                    size="small"
                    variant="outlined"
                />
                <p-menu #menu [model]="projectSettings" [popup]="true" />
            </div>
            <!-- -->
            <p-menu [model]="items" class="sidebar-menu" />
            <div class="grow"></div>
            <div class="p-4 ">
                <app-switch-workspaces />
            </div>
        </div>
    `
})
export class AppLayoutSidebar implements OnInit {
    private messageService = inject(MessageService);
    projectSettings: MenuItem[];
    items: MenuItem[] | undefined;

    ngOnInit() {
        this.projectSettings = [
            { label: "Members", icon: "pi pi-cog" },
            { label: "Permissions", icon: "pi pi-cog" },
            { label: "Integrations", icon: "pi pi-cog" },
            { label: "Webhooks", icon: "pi pi-cog" },
            {
                separator: true
            },
            { label: "Settings", icon: "pi pi-cog" }
        ];

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
