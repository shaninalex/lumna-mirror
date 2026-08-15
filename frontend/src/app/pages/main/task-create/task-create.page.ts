import { Component, DestroyRef, inject, signal } from '@angular/core';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { form, FormField, required } from '@angular/forms/signals';
import { MainLayout } from '@core/layout';
import { selectCurrentProjectId } from '@entities/project';
import { actionTaskCreate, actionTaskCreateSuccess, TaskCreateModel } from '@entities/task';
import { Actions, ofType } from '@ngrx/effects';
import { Store } from '@ngrx/store';
import { filter, tap } from 'rxjs';

@Component({
    selector: 'lu-task-create',
    imports: [MainLayout, FormField],
    templateUrl: './task-create.page.html',
})
export class TaskCreatePage {
    private store = inject(Store);
    private actions$ = inject(Actions);
    private destroyRef = inject(DestroyRef);

    taskFormModel = signal<TaskCreateModel>({
        title: '',
        project_id: 0,
    });

    taskForm = form(this.taskFormModel, (schemaPath) => {
        required(schemaPath.title);
    });

    constructor() {
        this.store
            .select(selectCurrentProjectId)
            .pipe(
                takeUntilDestroyed(this.destroyRef),
                filter((projectId) => projectId !== null),
                tap((projectId) => {
                    this.taskFormModel.set({
                        ...this.taskFormModel(),
                        project_id: projectId,
                    });
                }),
            )
            .subscribe();

        this.actions$
            .pipe(
                ofType(actionTaskCreateSuccess),
                takeUntilDestroyed(this.destroyRef),
                tap(() => {
                    this.taskFormModel.set({
                        ...this.taskFormModel(),
                        title: '',
                    });
                    this.taskForm().reset();
                }),
            )
            .subscribe();
    }

    onSubmit(event: Event): void {
        event.preventDefault();
        this.store.dispatch(actionTaskCreate({ data: this.taskFormModel() }));
    }
}
