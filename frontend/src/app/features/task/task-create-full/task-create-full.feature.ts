import { Component, DestroyRef, inject, signal } from '@angular/core';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { form, FormField, required } from '@angular/forms/signals';
import { Actions, ofType } from '@ngrx/effects';
import { Store } from '@ngrx/store';
import { filter, tap } from 'rxjs';

import { selectCurrentProjectId } from '@entities/project';
import type { TaskCreateModel} from '@entities/task';
import { actionTaskCreateSuccess, actionTaskCreate } from '@entities/task';

@Component({
    selector: 'lu-task-create-full-feature',
    imports: [FormField],
    templateUrl: './task-create-full.feature.html',
})
export class TaskCreateFullFeature {
    private store = inject(Store);
    private actions$ = inject(Actions);
    private destroyRef = inject(DestroyRef);

    taskFormModel = signal<TaskCreateModel>({
        title: '',
        project_id: 0,
        body: '',
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
                        body: '',
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
