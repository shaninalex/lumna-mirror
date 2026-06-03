import { Component, DestroyRef, inject, signal } from "@angular/core";
import { form, FormField, required } from "@angular/forms/signals";
import { takeUntilDestroyed } from "@angular/core/rxjs-interop";

import { Store } from "@ngrx/store";
import { Actions, ofType } from "@ngrx/effects";
import { filter, tap } from "rxjs/operators";

import { CardModule } from "primeng/card";
import { FormsModule } from "@angular/forms";
import { ButtonModule } from "primeng/button";
import { InputTextModule } from "primeng/inputtext";

import { selectProjectCurrentId } from "@entities/project";
import {
    actionTaskCreate,
    actionTaskCreateSuccess,
    TaskCreateModel
} from "@entities/task";

@Component({
    selector: "app-task-inline-form-feature",
    imports: [
        CardModule,
        FormsModule,
        ButtonModule,
        InputTextModule,
        FormField
    ],
    template: `
        <form (submit)="onSubmit($event)" class="flex gap-2">
            <input
                pInputText
                id="sprintTitle"
                placeholder="Task title"
                class="grow"
                [class.p-invalid]="
                    taskForm.title().touched() && taskForm.title().invalid()
                "
                [formField]="taskForm.title"
            />
            <p-button label="Create" type="submit" />
        </form>
    `
})
export class TaskInlineFormFeature {
    private store = inject(Store);
    private actions$ = inject(Actions);
    private destroyRef = inject(DestroyRef);

    taskFormModel = signal<TaskCreateModel>({
        title: "",
        project_id: 0
    });

    taskForm = form(this.taskFormModel, (schemaPath) => {
        required(schemaPath.title);
    });

    constructor() {
        this.store
            .select(selectProjectCurrentId)
            .pipe(
                takeUntilDestroyed(this.destroyRef),
                filter((projectId) => projectId !== undefined),
                tap((projectId) => {
                    this.taskFormModel.set({
                        ...this.taskFormModel(),
                        project_id: projectId
                    });
                })
            )
            .subscribe();

        this.actions$
            .pipe(
                ofType(actionTaskCreateSuccess),
                takeUntilDestroyed(this.destroyRef),
                tap(() => {
                    this.taskFormModel.set({
                        ...this.taskFormModel(),
                        title: ""
                    });
                    this.taskForm().reset();
                })
            )
            .subscribe();
    }

    onSubmit(event: Event): void {
        event.preventDefault();
        this.store.dispatch(actionTaskCreate({ data: this.taskFormModel() }));
    }
}
