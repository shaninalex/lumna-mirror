import {Component, inject, Input, OnDestroy, OnInit} from '@angular/core';
import {Task, TaskService} from '@client/entities/task';
import {MatInputModule} from '@angular/material/input';
import {ActivityListComponent} from '@client/entities/task/ui/activity-list';
import {TaskEditDetailsComponent} from '@client/features/task/task-edit-details/task-edit-details.component';
import {Subscription, tap} from 'rxjs';

@Component({
    selector: 'ts-task-detail-view',
    imports: [
        MatInputModule,
        ActivityListComponent,
        TaskEditDetailsComponent,
    ],
    template: `
        @if (!task) {
            loading...
        } @else {
            <div class="mb-4">
                <h4>{{ task.title }}</h4>
            </div>
            <div class="grid grid-cols-3 gap-4">
                <div class="col-span-2">
                    <ts-task-edit-details [task]="task" (update)="onUpdate($event)"/>
                </div>
                <div>
                    <ts-activity-list [taskID]="task.id"/>
                </div>
            </div>
        }
    `
})
export class TaskDetailViewComponent implements OnInit, OnDestroy {
    @Input() taskCode: string;
    @Input() projectCode: string;
    private _sub = new Subscription();
    api = inject(TaskService);
    task: Task;

    onUpdate(data: any) {
        this.api.Update(this.projectCode, this.taskCode, data).subscribe(result => {
            console.log(result)
        })
    }

    ngOnInit() {
        this._sub.add(
            this.api.Detail(this.projectCode, this.taskCode).pipe(
                tap(data => this.task = data)
            ).subscribe()
        )
    }

    ngOnDestroy() {
        this._sub.unsubscribe()
    }
}
