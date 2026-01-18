import { Component, input } from '@angular/core';
import { TaskModel } from '@entities/task';
import { DatePipe } from '@angular/common';

@Component({
    selector: 'app-task-card',
    imports: [DatePipe],
    templateUrl: './task-card.html',
})
export class TaskCard {
    task = input.required<TaskModel>();
}
