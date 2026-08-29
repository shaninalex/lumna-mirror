import { Component, input } from '@angular/core';
import { TaskDetailViewView } from '@features';
import { MainLayout } from '@core/layout';

@Component({
    selector: 'lu-task-detail-page',
    imports: [MainLayout, TaskDetailViewView],
    template: `
        <lu-main-layout>
            <div class="container-fluid">
                <lu-task-detail-view [taskId]="taskId()" />
            </div>
        </lu-main-layout>
    `,
})
export class TaskDetailPage {
    taskId = input.required({ transform: (id: string) => Number(id) });
}
