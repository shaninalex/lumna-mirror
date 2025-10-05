import {Component, Input} from '@angular/core';
import {Task} from '@client/entities/task';
import {DatePipe} from '@angular/common';
import {RouterLink} from '@angular/router';
import {MatCardModule} from '@angular/material/card';
import {MatIconModule} from '@angular/material/icon';

@Component({
    selector: "fr-task-card",
    imports: [
        DatePipe,
        RouterLink,
        MatCardModule,
        MatIconModule,
    ],
    template: `
        <mat-card appearance="outlined">
            <mat-card-header>
                <mat-card-title>
                    <a [routerLink]="['/projects', projectKey, task.code]">
                        {{ task.title }}
                        @if (task.completed) {
                            <mat-icon class="text-green-500">check_circle</mat-icon>
                        }
                    </a>
                </mat-card-title>
            </mat-card-header>
            <mat-card-content>
                <div class="flex items-center gap-2">
                    <div class="text-sm">
                        {{ task.created_at | date:"EEE, MMM d, HH:mm:ss" }} <br>
                        <small class="text-gray-500 dark:text-gray-300">{{ task.code }}</small>
                    </div>
                    <div class="ms-auto">
                        <img src="/img/1.png" class="rounded-full w-6">
                    </div>
                </div>
            </mat-card-content>
        </mat-card>
    `,
})
export class TaskCardComponent {
    @Input() task: Task;
    @Input() projectKey: string;
}
