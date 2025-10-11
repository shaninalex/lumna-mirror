import {Component, inject} from '@angular/core';
import {filter, map, Observable} from 'rxjs';
import {Task} from "@client/entities/task";
import {ActivatedRoute} from '@angular/router';
import {AsyncPipe} from '@angular/common';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {TaskDetailFeatureComponent} from '@client/features/task';

@Component({
    selector: "lu-task-detail-modal",
    imports: [
        AsyncPipe,
        FormsModule,
        ReactiveFormsModule,
        TaskDetailFeatureComponent
    ],
    template: `
        <div class="card">
            @if (task$ | async; as task) {
                <lu-task-detail-feature [task]="task" />
            }
        </div>
    `
})
export class TaskDetailPageComponent {
    private route = inject(ActivatedRoute)

    task$: Observable<Task> = this.route.data.pipe(
        map(data => data["task"]),
        filter(task => !!task),
    )
}
