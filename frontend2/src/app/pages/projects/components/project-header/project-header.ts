import { Component, Input } from "@angular/core";
import { ProjectModel } from "@entities/project";
import { Observable } from "rxjs";
import { AsyncPipe } from "@angular/common";
import { MegaMenuItem } from "primeng/api";
import { MegaMenuModule } from "primeng/megamenu";
import { ButtonModule } from "primeng/button";
import { MenuModule } from "primeng/menu";

@Component({
    selector: "app-project-header",
    imports: [AsyncPipe, MegaMenuModule, ButtonModule, MenuModule],
    template: `
        @if (project | async; as project) {
            <div class="p-4 border-b  border-slate-200">
                <div class="mb-3 flex items-center gap-4">
                    <div class="flex gap-2 items-center">
                        <img
                            src="/img/project.svg"
                            class="rounded"
                            style="width: 25px"
                        />
                        <div class="font-bold text-lg">
                            {{ project.title }}
                        </div>
                    </div>
                    <div class="grow"></div>
                    <div class="card flex justify-center">
                        <p-menu
                            #menu
                            [model]="projectSettings"
                            [popup]="true"
                        />
                        <p-button
                            (click)="menu.toggle($event)"
                            icon="pi pi-ellipsis-v"
                            outlined
                            size="small"
                        />
                    </div>
                </div>

                <p-megamenu [model]="items" />
            </div>
        }
    `
})
export class ProjectHeader {
    @Input() project: Observable<ProjectModel>;

    items: MegaMenuItem[] = [
        {
            label: "Summary",
            icon: "pi pi-file",
            routerLink: "summary"
        },
        {
            label: "Backlog",
            icon: "pi pi-box",
            routerLink: "backlog"
        },
        {
            label: "Board",
            icon: "pi pi-objects-column",
            routerLink: "board"
        }
    ];

    projectSettings: MegaMenuItem[] = [
        {
            label: "Project Settings",
            icon: "pi pi-cog",
            routerLink: "summary"
        }
    ];
}
