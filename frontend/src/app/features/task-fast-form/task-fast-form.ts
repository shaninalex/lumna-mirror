import { Component, effect, inject, input, signal } from '@angular/core';
import { ListModel } from '@entities/list';
import { FormField, form, required } from '@angular/forms/signals';
import { taskEvents, TaskPayloadModel } from '@entities/task';
import { Dispatcher, Events } from '@ngrx/signals/events';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';

@Component({
    selector: 'app-task-fast-form-feature',
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
export class TaskFastFormFeature {
    private readonly dispatcher = inject(Dispatcher);
    private readonly events = inject(Events);
    list = input.required<ListModel>();
    task_count = input.required<number>();
    openedForm = signal<boolean>(false);
    loading = signal(false);
    errors = signal<string[]>([]);
    taskFormModel = signal<TaskPayloadModel>({
        title: '',
        list_id: '',
        order: 0,
    });

    constructor() {
        this.events
            .on(taskEvents.setTask)
            .pipe(takeUntilDestroyed())
            .subscribe(() => {
                this.loading.set(false);
                this.openedForm.set(false);
                this.errors.set([]);
                this.taskForm().value.set({
                    title: '',
                    list_id: '',
                    order: 0,
                });
                this.errors.set([]);
            });

        this.events
            .on(taskEvents.failed)
            .pipe(takeUntilDestroyed())
            .subscribe((data) => {
                this.errors.set([data.payload.toString()]);
                this.loading.set(false);
            });

        effect(() => {
            if (!this.openedForm()) {
                this.taskForm().value.set({
                    title: '',
                    list_id: '',
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
        formData.list_id = this.list().id;
        formData.order = this.task_count();
        this.dispatcher.dispatch(
            taskEvents.create({
                board_id: this.list().board_id,
                data: formData,
            }),
        );
    }
}
