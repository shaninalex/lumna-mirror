import {Component} from '@angular/core';

@Component({
    selector: 'app-task-change-status',
    template: `
        <select class="form-select form-select-sm">
            <option selected>Change status</option>
            <option value="1">Done</option>
            <option value="2">Pending</option>
            <option value="3">Rejected</option>
        </select>
    `
})
export class TaskChangeStatusComponent {}
