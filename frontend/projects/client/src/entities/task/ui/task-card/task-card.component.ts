import {Component, inject, Input} from '@angular/core';
import {MatCardModule} from '@angular/material/card';
import {Task} from '@client/entities/task';
import {DatePipe} from '@angular/common';
import {MatButtonModule} from '@angular/material/button';
import {RouterLink} from '@angular/router';
import {FormsModule} from '@angular/forms';

@Component({
    selector: "fr-task-card",
    imports: [MatCardModule, DatePipe, MatButtonModule, RouterLink, FormsModule],
    template: `
        <mat-card appearance="outlined">
            <mat-card-header>
                <a [routerLink]="['/projects', projectKey, task.code]" class="cursor-pointer text-left">
                    {{ task.title }}
                </a>
            </mat-card-header>
            <mat-card-footer class="p-4 flex items-center gap-2">
                <div class="text-sm">
                    {{ task.created_at | date:"EEE, MMM d, HH:mm:ss" }} <br>
                    <small class="text-slate-500">{{ task.code }}</small>
                </div>
                <div class="ms-auto">
                    <img src="/img/1.png" class="rounded-full w-6">
                </div>
            </mat-card-footer>
        </mat-card>
    `,
})
export class TaskCardComponent {
    @Input() task: Task;
    @Input() projectKey: string;
}
