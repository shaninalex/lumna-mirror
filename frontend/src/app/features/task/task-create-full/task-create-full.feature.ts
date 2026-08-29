import { Component, DestroyRef, inject, signal } from '@angular/core';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { form, FormField, required } from '@angular/forms/signals';
import { Actions, ofType } from '@ngrx/effects';
import { Store } from '@ngrx/store';
import { actionTask, type TaskCreateModel} from '@entities/task';
import { ActivatedRoute, Router } from '@angular/router';
import { tap } from 'rxjs';

@Component({
    selector: 'lu-task-create-full-feature',
    imports: [FormField],
    templateUrl: './task-create-full.feature.html',
})
export class TaskCreateFullFeature {
    private store = inject(Store);
    private actions$ = inject(Actions);
    private destroyRef = inject(DestroyRef);
    private activatedRoute = inject(ActivatedRoute)
    private router = inject(Router);

    taskFormModel = signal<TaskCreateModel>({
        title: '',
        body: '',
        position: 0,
        column_id: 0,
        board_id: 0,
        project_id: 0, // TODO: get the value!
    });

    taskForm = form(this.taskFormModel, (schemaPath) => {
        required(schemaPath.title);
    });

    constructor() {
        this.actions$
            .pipe(
                ofType(actionTask.createSuccess),
                takeUntilDestroyed(this.destroyRef),
                tap(() => {
                    this.taskFormModel.set({
                        ...this.taskFormModel(),
                        title: '',
                        body: '',
                    });
                    this.taskForm().reset();
                    if ("return_to" in this.activatedRoute.snapshot.queryParams) {
                        this.router.navigate([this.activatedRoute.snapshot.queryParams["return_to"]])
                    }
                }),
            )
            .subscribe();
    }

    onSubmit(event: Event): void {
        event.preventDefault();
        this.store.dispatch(actionTask.create({ data: this.taskFormModel() }));
    }
}
