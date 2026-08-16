import { Component } from '@angular/core';
import { MainLayout } from '@core/layout';
import { TaskCreateFullFeature } from '@features/task';

@Component({
    selector: 'lu-task-create-page',
    imports: [MainLayout, TaskCreateFullFeature],
    templateUrl: './task-create.page.html',
})
export class TaskCreatePage {}
