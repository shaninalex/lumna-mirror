import { Component, inject } from "@angular/core";
import { MenuModule } from "primeng/menu";
import { Store } from "@ngrx/store";
import { selectCurrentProjectList } from "@entities/project";
import { AsyncPipe } from "@angular/common";
import { ButtonModule } from "primeng/button";
import { RouterLink } from "@angular/router";
import { ButtonGroupModule } from "primeng/buttongroup";
import { map } from "rxjs";
import { filter } from "rxjs/operators";

@Component({
    selector: "app-project-list",
    imports: [
        MenuModule,
        AsyncPipe,
        ButtonModule,
        RouterLink,
        ButtonGroupModule
    ],
    template: `
        @if (items$ | async; as items) {
            <p-menu [model]="items" class="sidebar-menu" />
        }
        <div class="px-4">
            <p-buttonGroup>
                <p-button
                    label="Create"
                    icon="pi pi-plus"
                    size="small"
                    variant="outlined"
                    routerLink="projects/create"
                />
                <p-button
                    label="View all projects"
                    size="small"
                    variant="outlined"
                    routerLink="projects"
                />
            </p-buttonGroup>
        </div>
    `
})
export class ProjectList {
    private store = inject(Store);
    items$ = this.store.select(selectCurrentProjectList).pipe(
        filter((projects) => !!projects),
        map((projects) =>
            projects.map((project) => ({
                label: project.title,
                icon: "pi pi-th-large",
                routerLink: project.appLink
            }))
        )
    );
}
