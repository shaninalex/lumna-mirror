import { Component, inject } from "@angular/core";
import { MenuModule } from "primeng/menu";
import { Store } from "@ngrx/store";
import { selectCurrentProjectList } from "@entities/project";
import { combineLatest, map } from "rxjs";
import { AsyncPipe } from "@angular/common";
import { ButtonModule } from "primeng/button";
import { selectWorkspaceCurrentWorkspaceSlug } from "@entities/workspace";

@Component({
    selector: "app-project-list",
    imports: [MenuModule, AsyncPipe, ButtonModule],
    template: `
        @if (items$ | async; as items) {
            <p-menu [model]="items" class="sidebar-menu" />
        }

        <div class="px-4">
            <p-button label="Create project" size="small" variant="outlined" />
        </div>
    `
})
export class ProjectList {
    private store = inject(Store);

    items$ = combineLatest([
        this.store.select(selectWorkspaceCurrentWorkspaceSlug),
        this.store.select(selectCurrentProjectList)
    ]).pipe(
        map(([workspaceSlug, projects]) => {
            if (!projects || !workspaceSlug) {
                return [];
            }
            return projects.map((project) => ({
                label: project.title,
                icon: "pi pi-th-large",
                routerLink: `/app/${workspaceSlug}/project/${project.id}`
            }));
        })
    );
}
