import { Component, inject } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { filter, map, Observable, tap } from 'rxjs';
import { TaskModel } from '@entities/task';
import { UiService } from '@shared/ui';
import { AsyncPipe } from '@angular/common';

@Component({
    selector: 'app-task-detail-page',
    template: `
        @if(task$ | async; as task) {
            <h1>{{ task.title }}</h1>
        }
    `,
    imports: [AsyncPipe],
})
export class TaskDetailComponent {
    private route = inject(ActivatedRoute);
    private ui = inject(UiService);

    task$: Observable<TaskModel> = this.route.data.pipe(
        filter((data) => !!data['task']),
        map((data) => data['task'] as TaskModel),
        tap((task) => this.ui.setPageTitle(`Task: ${task.title}`)),
    );
}
