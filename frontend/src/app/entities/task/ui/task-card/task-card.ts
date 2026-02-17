import { Component, input } from '@angular/core';
import { TaskModel } from '@entities/task';
import { DatePipe } from '@angular/common';
import { RouterLink } from '@angular/router';

@Component({
    selector: 'app-task-card',
    imports: [DatePipe, RouterLink],
    templateUrl: './task-card.html',
})
export class TaskCard {
    task = input.required<TaskModel>();
}
