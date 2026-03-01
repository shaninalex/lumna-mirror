import {Component, Input} from '@angular/core';

@Component({
    selector: 'app-task-activity',
    template: `
        <div class="d-flex align-items-center justify-content-between">
            <div class="fw-medium">
                <i class="fa-regular fa-message"></i>
                Comments and activity
            </div>
            <button type="button" class="btn btn-secondary btn-sm">Details</button>
        </div>
    `
})
export class TaskActivityComponent {
    @Input() task_id: number;
}
