import { Component } from '@angular/core';
import { MainLayout } from '@core/layout';
import { RouterLink } from "@angular/router";

@Component({
    selector: 'lu-task-page',
    imports: [MainLayout, RouterLink],
    templateUrl: './task.page.html',
})
export class TaskPage {}
