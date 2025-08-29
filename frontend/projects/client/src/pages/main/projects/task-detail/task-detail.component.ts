import {Component, inject} from '@angular/core';
import {TaskDetailViewComponent} from '@client/entities/task/ui/task-detail-view';
import {ActivatedRoute} from '@angular/router';
import {AsyncPipe} from '@angular/common';

@Component({
    selector: 'ts-task-detail',
    imports: [
        TaskDetailViewComponent,
        AsyncPipe,
    ],
    template: `
        @if (route.params | async; as params) {
            <ts-task-detail-view [projectCode]="params['projectKey']" [taskCode]="params['taskCode']" />
        }
    `
})
export class TaskDetailComponent {
    // get task code and project key from url
    route = inject(ActivatedRoute)
}
