import {Component, Input} from '@angular/core';
import {Task} from '@client/entities/task';
import {DatePipe} from '@angular/common';
import {RouterLink} from '@angular/router';
import {FormsModule} from '@angular/forms';

@Component({
    selector: "fr-task-card",
    imports: [DatePipe, RouterLink, FormsModule],
    template: `
        <div class="card border border-base-300 bg-amber-100 dark:bg-stone-800">
            <div class="card-body">
                <a [routerLink]="['/projects', projectKey, task.code]" class="card-title cursor-pointer text-left">
                    {{ task.title }}
                </a>
                <div class="flex items-center gap-2">
                    <div class="text-sm">
                        {{ task.created_at | date:"EEE, MMM d, HH:mm:ss" }} <br>
                        <small class="text-gray-500 dark:text-gray-300">{{ task.code }}</small>
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
