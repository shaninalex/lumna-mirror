import { Component, OnInit } from "@angular/core";
import { MenuModule } from "primeng/menu";
import { MenuItem } from "primeng/api";

@Component({
    selector: "app-project-list",
    imports: [MenuModule],
    template: ` <p-menu [model]="items" class="sidebar-menu" /> `
})
export class ProjectList implements OnInit {
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
