import {Component, inject, Input, OnInit} from '@angular/core';
import {Store} from '@ngrx/store';
import {AppState} from '@client/shared/store';
import {Task, TaskDetailInput, TaskPatchAction} from '@client/entities/task';
import {FormControl, FormGroup, FormsModule, ReactiveFormsModule, Validators} from '@angular/forms';
import {DatePipe} from '@angular/common';
import {MoveTaskComponent} from './move-task';

@Component({
    selector: "fr-task-detail-feature",
    imports: [
        DatePipe,
        FormsModule,
        ReactiveFormsModule,
        MoveTaskComponent
    ],
    template: `
        <div class="mb-4">
            <lu-move-task [task]="task" />
        </div>
        <form [formGroup]="form" (ngSubmit)="onSubmit()">
            Title
            <div class="mb-4 card-title">
                <input class="input w-full" formControlName="title" />
                <div class="text-xs">Created: {{ task.created_at | date }}</div>
            </div>

            <div class="mb-4">
                <div class="font-bold text-sm">Description</div>
                <textarea rows="8" class="input" formControlName="description"></textarea>
            </div>

            <div class="mb-4">
                <button [disabled]="!form.valid" class="btn btn-primary" type="submit">Save</button>
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
