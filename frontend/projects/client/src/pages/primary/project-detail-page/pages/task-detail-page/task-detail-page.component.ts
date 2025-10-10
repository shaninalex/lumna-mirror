import {Component, inject, OnDestroy, OnInit} from '@angular/core';
import {filter, map, Observable, tap} from 'rxjs';
import {Task, TaskDetailInput, TaskPatchAction, TaskService} from "@client/entities/task";
import {ActivatedRoute} from '@angular/router';
import {AsyncPipe, DatePipe} from '@angular/common';
import {FormControl, FormGroup, FormsModule, ReactiveFormsModule, Validators} from '@angular/forms';
import {Store} from '@ngrx/store';
import {AppState} from '@client/shared/store';

@Component({
    selector: "fr-task-detail-modal",
    imports: [
        AsyncPipe,
        DatePipe,
        FormsModule,
        ReactiveFormsModule
    ],
    template: `
        <div class="card">
            @if (task$ | async; as task) {
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
            }
        </div>
    `
})
export class TaskDetailPageComponent {
    private route = inject(ActivatedRoute)
    private store = inject(Store<AppState>)
    private task: Task

    form: FormGroup = new FormGroup({
        title: new FormControl('', Validators.required),
        description: new FormControl('', Validators.required),
    })

    task$: Observable<Task> = this.route.data.pipe(
        map(data => data["task"]),
        filter(task => !!task),
        tap(task => {
            this.task = task
            this.form.setValue({
                title: task.title,
                description: task.description,
            })
        })
    )

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
