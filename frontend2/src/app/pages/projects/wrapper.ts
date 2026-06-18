import { Component, inject } from "@angular/core";
import { ActivatedRoute, RouterOutlet } from "@angular/router";
import { ProjectHeader } from "./components";
import { Store } from "@ngrx/store";
import { filter, switchMap, tap } from "rxjs/operators";
import { map } from "rxjs";
import {
    actionProjectSetCurrentProjectId,
    selectProject
} from "@entities/project";
import { actionTaskGetList } from "@entities/task";
import { actionSprintGetList } from "@entities/sprint";

@Component({
    selector: "app-project-wrapper",
    imports: [RouterOutlet, ProjectHeader],
    template: `
        <app-project-header [project]="project$" />
        <router-outlet />
    `
})
export class ProjectWrapper {
    private route = inject(ActivatedRoute);
    private store = inject(Store);

    project$ = this.route.params.pipe(
        filter((params) => params["project-id"]),
        map((params) => parseInt(params["project-id"])),
        switchMap((id) => this.store.select(selectProject(id))),
        filter((project) => project !== undefined),
        tap((project) => {
            this.store.dispatch(
                actionProjectSetCurrentProjectId({ project_id: project.id })
            );
            this.store.dispatch(
                actionTaskGetList({
                    query: {
                        project_id: project.id
                    }
                })
            );
            this.store.dispatch(
                actionSprintGetList({
                    query: {
                        project_id: project.id
                    }
                })
            );
        })
    );
}
