import { Component, inject } from "@angular/core";
import { actionSprintCreate, SprintWideList } from "@entities/sprint";
import { SprintSearchFormBarFeature, TaskInlineFormFeature } from "@features";
import { Store } from "@ngrx/store";
import { AsyncPipe } from "@angular/common";
import { CardModule } from "primeng/card";
import { selectTasksByCurrentProject } from "@entities/task";
import { Listbox } from "primeng/listbox";
import { FormsModule } from "@angular/forms";
import { Button } from "primeng/button";

@Component({
    selector: "app-backlog-page",
    imports: [
        SprintWideList,
        SprintSearchFormBarFeature,
        AsyncPipe,
        TaskInlineFormFeature,
        CardModule,
        Listbox,
        FormsModule,
        Button
    ],
    template: `
        <div class="min-h-screen bg-slate-100 text-slate-900">
            <main class="mx-auto p-4">
                <app-sprint-search-form-bar-feature />

                <app-sprint-wide-list />

                <p-card>
                    <div class="flex justify-between items-center">
                        <div>
                            <div class="font-bold">Backlog</div>
                            @if (tasks$ | async; as tasks) {
                                <div class="text-sm text-slate-500">
                                    {{ tasks.length }} tasks in backlog
                                </div>
                            }
                        </div>
                        <div>
                            <p-button
                                label="Create sprint"
                                variant="outlined"
                                severity="secondary"
                                type="button"
                                (click)="createSprint($event)"
                            />
                        </div>
                    </div>

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

    tasks$ = this.store.select(selectTasksByCurrentProject);
    selectedTask?: string;

    onChange(event: any) {
        console.log(event);
    }

    createSprint(event: any): void {
        this.store.dispatch(
            actionSprintCreate({
                data: {
                    title: "New sprint"
                }
            })
        );
    }
}
