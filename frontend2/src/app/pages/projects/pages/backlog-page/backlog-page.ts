import { Component, inject } from "@angular/core";
import { selectSprintByProject, SprintWideListItem } from "@entities/sprint";
import { SprintCreateFeature, SprintSearchFormBarFeature, TaskInlineFormFeature } from "@features";
import { Store } from "@ngrx/store";
import { selectProjectCurrentId } from "@entities/project";
import { AsyncPipe } from "@angular/common";
import { filter, switchMap } from "rxjs/operators";
import { CardModule } from "primeng/card";

@Component({
    selector: "app-backlog-page",
    imports: [
        SprintWideListItem,
        SprintSearchFormBarFeature,
        SprintCreateFeature,
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
                    <div>
                        <!-- tasks without sprint -->
                    </div>
                    <app-task-inline-form-feature />
                </p-card>
                <app-sprint-create-feature />
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
}
