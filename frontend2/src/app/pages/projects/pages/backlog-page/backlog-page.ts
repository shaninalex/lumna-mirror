import { Component, inject } from "@angular/core";
import { selectSprintByProject, SprintWideListItem } from "@entities/sprint";
import { SprintSearchFormBarFeature, TaskInlineFormFeature } from "@features";
import { Store } from "@ngrx/store";
import { selectProjectCurrentId } from "@entities/project";
import { AsyncPipe } from "@angular/common";
import { filter, switchMap } from "rxjs/operators";
import { CardModule } from "primeng/card";
import { selectTasksByCurrentProject } from "@entities/task";

@Component({
    selector: "app-backlog-page",
    imports: [
        SprintWideListItem,
        SprintSearchFormBarFeature,
        AsyncPipe,
        TaskInlineFormFeature,
        CardModule
    ],
    template: `
        <div class="min-h-screen bg-slate-100 text-slate-900">
            <main class="mx-auto p-4">
                <app-sprint-search-form-bar-feature />
                @if (sprints$ | async; as sprints) {
                    @for (sprint of sprints; track sprint.id) {
                        <app-sprint-wide-list-item [sprint]="sprint" />
                    }
                }

                <p-card>
                    @if (tasks$ | async; as tasks) {
                        <div>
                            @for (task of tasks; track task.id) {
                                <div>{{ task.title }}</div>
                            }
                        </div>
                    }
                    <app-task-inline-form-feature />
                </p-card>
            </main>
        </div>
    `
})
export class BacklogPage {
    private store = inject(Store);
    sprints$ = this.store.select(selectProjectCurrentId).pipe(
        filter((projectId) => projectId !== undefined),
        switchMap((projectId) =>
            this.store.select(selectSprintByProject(projectId))
        )
    );

    tasks$ = this.store.select(selectTasksByCurrentProject);
}
