import {Component, inject, Input, OnInit} from '@angular/core';
import {Store} from '@ngrx/store';
import {
    actionActivityGetList, ActivityListItemComponent,
    ActivityModel,
    selectActivity
} from '@entities/activity';
import {Observable} from 'rxjs';
import {AsyncPipe} from '@angular/common';

@Component({
    selector: 'app-task-activity',
    imports: [
        AsyncPipe,
        ActivityListItemComponent
    ],
    template: `
        <div class="d-flex align-items-center justify-content-between mb-3">
            <div class="fw-medium">
                <i class="fa-regular fa-message"></i>
                Comments and activity
            </div>
            <button type="button" class="btn btn-outline-secondary btn-sm">Details</button>
        </div>

        @if (activityLog | async; as activityLog) {
            <div class="d-flex flex-column gap-2">
                @for (activity of activityLog; track activity.id) {
                    <app-activity-list-item [activity]="activity" />
                }
            </div>
        }
    `
})
export class TaskActivityComponent implements OnInit {
    @Input() task_id: number;
    private store = inject(Store)
    activityLog: Observable<ActivityModel[]>

    ngOnInit(): void {
        this.store.dispatch(actionActivityGetList({entity_id: this.task_id, entity_type: "task"}))
        this.activityLog = this.store.select(selectActivity(this.task_id, "task"))
    }
}
