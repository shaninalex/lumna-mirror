import { DatePipe } from '@angular/common';
import { Component, computed, inject, input } from '@angular/core';
import { RouterLink } from '@angular/router';
import { AppRoutes } from '@core/routes';
import { selectTask } from '@entities/task';
import { Store } from '@ngrx/store';
import { TimeAgoPipe } from '@shared/utils';

@Component({
    selector: 'lu-task-detail-view',
    imports: [RouterLink, DatePipe, TimeAgoPipe],
    templateUrl: './task-detail-view.view.html',
})
export class TaskDetailViewView {
    private store = inject(Store);
    readonly appRouter = inject(AppRoutes)
    taskId = input.required<number>();
    task = computed(() => this.store.selectSignal(selectTask(this.taskId()))());
}
