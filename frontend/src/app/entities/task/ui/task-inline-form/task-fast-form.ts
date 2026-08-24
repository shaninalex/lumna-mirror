import { Component, effect, inject, input, signal } from '@angular/core';
import { FormField, form, required } from '@angular/forms/signals';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { Store } from '@ngrx/store';
import { Actions, ofType } from '@ngrx/effects';
import { actionTask, type TaskCreateModel } from '@entities/task/model';

@Component({
    selector: 'lu-task-inline-form',
    imports: [FormField],
    template: `
        @if (openedForm()) {
            <form (submit)="submit($event)" class="w-full">
                <div class="mb-2">
                    <input
                        class="form-control form-control-sm w-full"
                        placeholder="Column name"
                        [formField]="taskForm.title"
                        [autofocus]="true"
                    />
                    @if (taskForm.title().touched() && taskForm.title().errors().length) {
                        @for (error of taskForm.title().errors(); track error) {
                            <div class="text-danger small">{{ error.message }}</div>
                        }
                    }
                </div>

                <div class="d-flex gap-2">
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
            <button class="btn btn-sm btn-outline-secondary" (click)="this.openedForm.set(true)">
                add task
            </button>
        }
    `,
})
export class TaskInlineForm {
    private store = inject(Store);
    private actions$ = inject(Actions);

    column_id = input.required<number>();
    task_count = input.required<number>();

    openedForm = signal<boolean>(false);
    loading = signal(false);
    errors = signal<string[]>([]);
    taskFormModel = signal<{ title: string }>({ title: '' });
    taskForm = form(this.taskFormModel, (schemaPath) => {
        required(schemaPath.title, { message: 'Name is required' });
    });

    constructor() {
        this.actions$
            .pipe(ofType(actionTask.createSuccess), takeUntilDestroyed())
            .subscribe(() => this._reset());

        this.actions$.pipe(ofType(actionTask.createFailed)).subscribe((data) => {
            this.errors.set([data.errors.toString()]);
            this.loading.set(false);
        });

        effect(() => {
            if (!this.openedForm()) {
                this.taskForm().value.set({ title: '' });
                this.errors.set([]);
            }
        });
    }

    submit(event: Event) {
        event.preventDefault();
        const formData = this.taskFormModel();
        const data: TaskCreateModel = {
            title: formData.title,
            body: '',
            order: this.task_count(),
            status_id: this.column_id(),
        };
        this.store.dispatch(actionTask.create({ data }));
        this._reset();
    }

    private _reset() {
        this.loading.set(false);
        this.openedForm.set(false);
        this.errors.set([]);
        this.taskForm().reset();
    }
}
