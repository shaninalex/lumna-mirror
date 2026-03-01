import { Component, Input, input } from '@angular/core';
import { TaskModel } from '@entities/task';
import { DatePipe } from '@angular/common';
import { RouterLink } from '@angular/router';

@Component({
    selector: 'app-task-card',
    imports: [DatePipe, RouterLink],
    templateUrl: './task-card.html',
})
export class TaskCard {
    @Input() board_id: number;
    @Input() task: TaskModel;

    detailUrl(): string {
        return `/projects/${this.task.project_id}/boards/${this.board_id}/task/${this.task.id}`;
    }
}
