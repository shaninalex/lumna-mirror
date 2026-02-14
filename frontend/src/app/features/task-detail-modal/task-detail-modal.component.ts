import { Component, Input } from '@angular/core';
import { TaskModel } from '@entities/task';

@Component({
    selector: 'app-task-detail-modal-feature',
    template: './task-detail-modal.component.html',
})
export class TaskDetailModalFeature {
    @Input() task_id: string;
}
