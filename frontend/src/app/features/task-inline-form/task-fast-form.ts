import { Component, effect, inject, input, signal } from '@angular/core';
import { FormField, form, required } from '@angular/forms/signals';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { Store } from '@ngrx/store';
import { Actions, ofType } from '@ngrx/effects';
import {
    TaskState,
    TaskPayloadModel,
    actionTaskUpsert,
    actionTaskFailed,
    actionTaskCreate,
} from '@entities/task';
import { ColumnModel } from '@entities/column';

@Component({
    selector: 'app-task-inline-form-feature',
    imports: [FormField],
    template: `
        @if (openedForm()) {
            <form (submit)="submit($event)" class="w-full">
                <div class="mb-2">
                    <input
                        class="input w-full"
                        placeholder="Column name"
                        [formField]="taskForm.title"
                        [autofocus]="true"
                    />
                    @if (taskForm.title().touched() && taskForm.title().errors().length) {
                        <ul class="error-list">
                            @for (error of taskForm.title().errors(); track error) {
                                <li class="text-red-500 text-sm">{{ error.message }}</li>
                            }
                        </ul>
                    }
                </div>

                <div class="flex gap-2">
                    <button
                        class="btn btn-sm btn-primary"
                        (click)="submit($event)"
                        [disabled]="taskForm().invalid()"
                    >
                        @if (loading()) {
                            Processing...
                        } @else {
                            Create
                        }
                    </button>
                    <button
                        type="button"
                        class="btn btn-sm btn-secondary"
                        (click)="this.openedForm.set(false)"
                    >
                        cancel
                    </button>
                </div>
            </form>
        } @else {
            <button class="btn btn-secondary btn-sm" (click)="this.openedForm.set(true)">
                add task
            </button>
        }
    `,
})
export class TaskInlineFormFeature {
    private store = inject(Store<TaskState>);
    private actions$ = inject(Actions);

    board_id = input.required<string>();
    project_id = input.required<string>();
    column_id = input.required<string>();
    task_count = input.required<number>();

    openedForm = signal<boolean>(false);
    loading = signal(false);
    errors = signal<string[]>([]);
    taskFormModel = signal<TaskPayloadModel>({
        title: '',
        column_id: '',
        project_id: '',
        order: 0,
    });

    constructor() {
        this.actions$
            .pipe(ofType(actionTaskUpsert), takeUntilDestroyed())
            .subscribe(() => this._reset());

        this.actions$.pipe(ofType(actionTaskFailed), takeUntilDestroyed()).subscribe((data) => {
            this.errors.set([data.error.toString()]);
            this.loading.set(false);
        });

        effect(() => {
            if (!this.openedForm()) {
                this.taskForm().value.set({
                    title: '',
                    column_id: '',
                    project_id: '',
                    order: 0,
                });
                this.errors.set([]);
            }
        });
    }

    taskForm = form(this.taskFormModel, (schemaPath) => {
        required(schemaPath.title, { message: 'Name is required' });
    });

    submit(event: Event) {
        event.preventDefault();
        const formData = this.taskFormModel();
        const payload = {
            title: formData.title,
            order: formData.order,
            column_id: this.column_id(),
            project_id: this.board_id(),
        };
        this.store.dispatch(
            actionTaskCreate({
                data: payload,
            }),
        );
    }

    private _reset() {
        this.loading.set(false);
        this.openedForm.set(false);
        this.errors.set([]);
        this.taskForm().value.set({
            title: '',
            column_id: '',
            project_id: '',
            order: 0,
        });
        this.errors.set([]);
    }
}
