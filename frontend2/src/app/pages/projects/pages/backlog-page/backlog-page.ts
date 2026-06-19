import { Component, inject } from "@angular/core";
import { selectSprintByProject, SprintWideListItem } from "@entities/sprint";
import { SprintSearchFormBarFeature, TaskInlineFormFeature } from "@features";
import { Store } from "@ngrx/store";
import { selectProjectCurrentId } from "@entities/project";
import { AsyncPipe } from "@angular/common";
import { filter, switchMap } from "rxjs/operators";
import { CardModule } from "primeng/card";
import { selectTasksByCurrentProject } from "@entities/task";
import { Listbox } from "primeng/listbox";
import { FormsModule } from "@angular/forms";

@Component({
    selector: "app-backlog-page",
    imports: [
        SprintWideListItem,
        SprintSearchFormBarFeature,
        AsyncPipe,
        TaskInlineFormFeature,
        CardModule,
        Listbox,
        FormsModule
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
                        <p-listbox
                            [options]="tasks"
                            [(ngModel)]="selectedTask"
                            (onChange)="onChange($event)"
                            [striped]="true"
                            [dragdrop]="true"
                            [scrollHeight]="''"
                            optionLabel="title"
                            optionValue="code"
                            dataKey="code"
                            class="mb-4"
                            style="border: none!important; box-shadow: none!important;"
                        >
                            <ng-template #item let-task let-selected="selected">
                                <div
                                    class="flex w-full items-center gap-3 py-1"
                                >
                                    <i
                                        class="pi"
                                        [class.pi-check-circle]="task.done"
                                        [class.pi-circle]="!task.done"
                                        [class.text-green-500]="task.done"
                                        [class.text-slate-400]="!task.done"
                                    ></i>
                                    <span
                                        class="font-mono text-slate-500 shrink-0"
                                    >
                                        {{ task.code }}
                                    </span>
                                    <span
                                        class="flex-1 truncate"
                                        [class.line-through]="task.done"
                                    >
                                        {{ task.title }}
                                    </span>
                                    <span
                                        class="rounded px-2 py-0.5 text-[11px] font-semibold uppercase shrink-0"
                                        [class.bg-green-100]="task.done"
                                        [class.text-green-700]="task.done"
                                        [class.bg-slate-200]="!task.done"
                                        [class.text-slate-600]="!task.done"
                                    >
                                        {{ task.done ? "Done" : "To do" }}
                                    </span>
                                </div>
                            </ng-template>
                        </p-listbox>
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
    selectedTask?: string;

    onChange(event: any) {
        console.log(event);
    }
}
