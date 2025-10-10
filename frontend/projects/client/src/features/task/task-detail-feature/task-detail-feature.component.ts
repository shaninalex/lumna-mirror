import {Component, inject, Input, OnInit} from '@angular/core';
import {Store} from '@ngrx/store';
import {AppState} from '@client/shared/store';
import {Task, TaskDetailInput, TaskPatchAction} from '@client/entities/task';
import {FormControl, FormGroup, FormsModule, ReactiveFormsModule, Validators} from '@angular/forms';
import {AsyncPipe, DatePipe} from '@angular/common';

@Component({
    selector: "fr-task-detail-feature",
    imports: [
        DatePipe,
        FormsModule,
        ReactiveFormsModule
    ],
    template: `
        <form [formGroup]="form" (ngSubmit)="onSubmit()">
            <div class="flex items-start justify-between gap-4">
                <div class="mb-4 flex-grow">
                    <div class="card-title">
                        <input class="input w-full" formControlName="title" />
                    </div>
                    <div class="text-xs">Created: {{ task.created_at | date }}</div>
                </div>
                <div>
                    <button [disabled]="!form.valid" class="btn btn-primary" type="submit">Save</button>
                </div>
            </div>

            <div>
                <div class="font-bold text-sm">Description</div>
                <textarea rows="8" class="input" formControlName="description"></textarea>
            </div>
        </form>
    `
})
export class TaskDetailFeatureComponent implements OnInit {
    @Input() task: Task
    private store = inject(Store<AppState>)

    form: FormGroup = new FormGroup({
        title: new FormControl('', Validators.required),
        description: new FormControl('', Validators.required),
    })

    ngOnInit() {
        this.form.setValue({
            title: this.task.title,
            description: this.task.description,
        })
    }

    onSubmit(): void {
        const payload: TaskDetailInput = {
            title: this.form.value["title"],
            completed: this.task.completed,
            description: this.form.value["description"],
            list_index: this.task.list_index,
            status_id: this.task.status_id,
        }
        this.store.dispatch(TaskPatchAction({taskId: this.task.id, payload }))
    }
}
