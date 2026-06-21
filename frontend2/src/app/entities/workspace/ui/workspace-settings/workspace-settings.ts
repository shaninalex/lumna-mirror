import { Component, OnInit } from "@angular/core";
import { Button } from "primeng/button";
import { Menu } from "primeng/menu";
import { MenuItem } from "primeng/api";

@Component({
    selector: "app-workspace-settings",
    imports: [Button, Menu],
    template: `
        <p-menu #menu [model]="items" [popup]="true" />
        <p-button
            (click)="menu.toggle($event)"
            label="Workspace settings"
            size="small"
            variant="outlined"
        />
    `
})
export class WorkspaceSettings implements OnInit {
    items: MenuItem[];

    ngOnInit() {
        this.items = [
            { label: "Members", icon: "pi pi-cog" },
            { label: "Permissions", icon: "pi pi-cog" },
            { label: "Integrations", icon: "pi pi-cog" },
            { label: "Webhooks", icon: "pi pi-cog" },
            {
                separator: true
            },
            { label: "Settings", icon: "pi pi-cog" }
        ];
    }
}
