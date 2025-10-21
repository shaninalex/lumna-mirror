import {Component, inject, Input, OnInit} from '@angular/core';
import {Store} from '@ngrx/store';
import {AppState} from '@client/shared/store';
import {Task, TaskChangeStatusAction} from '@client/entities/task';
import {selectProjectStatusList, Status} from '@client/entities/status';
import {Observable} from 'rxjs';
import {AsyncPipe} from '@angular/common';
import {FormsModule} from '@angular/forms';

@Component({
    selector: "lu-move-task",
    imports: [
        AsyncPipe,
        FormsModule
    ],
    template: `
        Move task:
        @if (statuses$ | async; as statuses) {
            <select [(ngModel)]="status_id" name="statuses" id="statuses" class="input" (change)="onChangeStatus()">
                @for (status of statuses; track status.id) {
                    <option [value]="status.id">
                        {{ status.title }}
                    </option>
                }
            </select>
        }
    `
})
export class MoveTaskComponent implements OnInit {
    @Input() task: Task;
    private store = inject(Store<AppState>)
    statuses$: Observable<Status[]>
    status_id: string

    ngOnInit() {
        this.statuses$ = this.store.select(selectProjectStatusList(this.task.project_id))
        this.status_id = this.task.status_id.toString();
    }

    onChangeStatus(): void {
        this.store.dispatch(TaskChangeStatusAction({
            taskId: this.task.id,
            payload: {
                from_status: this.task.status_id,
                to_status: parseInt(this.status_id),
                from_idx: 0,
                to_idx: this.task.list_index,
            }
        }))
    }
}
