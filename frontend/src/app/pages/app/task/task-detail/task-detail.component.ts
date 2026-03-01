import { Component, inject } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { filter, map, Observable, tap } from 'rxjs';
import { TaskModel } from '@entities/task';
import { UiService } from '@shared/ui';
import { AsyncPipe } from '@angular/common';
import {TaskDetailModalFeature} from '@features';

@Component({
    selector: 'app-task-detail-page',
    template: `
        <task-detail-modal-feature [task$]="task$" (onClose)="handleOnClose()"/>
    `,
    imports: [TaskDetailModalFeature],
})
export class TaskDetailComponent {
    private route = inject(ActivatedRoute);
    private router = inject(Router);
    private ui = inject(UiService);

    task$: Observable<TaskModel> = this.route.data.pipe(
        filter((data) => !!data['task']),
        map((data) => data['task'] as TaskModel),
        tap((task) => this.ui.setPageTitle(`Task: ${task.title}`)),
    );

    handleOnClose(): void {
        this.router.navigate(['../..'], { relativeTo: this.route });
    }
}
