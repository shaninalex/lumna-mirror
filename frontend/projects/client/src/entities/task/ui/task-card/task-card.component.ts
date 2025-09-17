import {Component, Input} from '@angular/core';
import {Task} from '@client/entities/task';
import {DatePipe} from '@angular/common';
import {MatButtonModule} from '@angular/material/button';
import {RouterLink} from '@angular/router';
import {FormsModule} from '@angular/forms';

@Component({
    selector: "fr-task-card",
    imports: [DatePipe, MatButtonModule, RouterLink, FormsModule],
    template: `
        <div class="card border">
            <div class="card-body">
                <a [routerLink]="['/projects', projectKey, task.code]" class="card-title cursor-pointer text-left">
                    {{ task.title }}
                </a>
                <div class="flex items-center gap-2">
                    <div class="text-sm">
                        {{ task.created_at | date:"EEE, MMM d, HH:mm:ss" }} <br>
                        <small class="text-slate-500 dark:text-slate-300">{{ task.code }}</small>
                    </div>
                    <div class="ms-auto">
                        <img src="/img/1.png" class="rounded-full w-6">
                    </div>
                </div>
            </div>
        </div>
    `,
})
export class TaskCardComponent {
    @Input() task: Task;
    @Input() projectKey: string;
}
