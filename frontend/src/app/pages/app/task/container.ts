import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';

@Component({
    selector: 'app-task-container',
    imports: [RouterOutlet],
    template: `<router-outlet />`,
})
export class TaskContainer {}
