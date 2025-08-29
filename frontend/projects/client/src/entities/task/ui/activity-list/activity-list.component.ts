import {Component, Input} from '@angular/core';

@Component({
    selector: "ts-activity-list",
    template: `
        <div>Task activity List</div>
        <code>{{ taskID }}</code>
    `
})
export class ActivityListComponent {
    @Input() taskID: string;
}
