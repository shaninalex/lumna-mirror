import { DatePipe } from '@angular/common';
import { Component, inject, Input } from '@angular/core';
import { RouterLink } from '@angular/router';
import type { TaskModel } from '@entities/task/model';
import { AppRoutes } from '@core';
import { standardTimeFormat } from '@shared/utils'

@Component({
    selector: 'lu-task-list-item',
    imports: [RouterLink, DatePipe],
    templateUrl: './task-list-item.component.html',
    host: {
        class: "list-group-item list-group-item-action",
    }
})
export class TaskListItemComponent {
    @Input() task: TaskModel;
    readonly appRoutes = inject(AppRoutes)
    standardTime = standardTimeFormat;
}
