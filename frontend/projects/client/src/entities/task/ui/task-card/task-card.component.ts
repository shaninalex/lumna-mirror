import {Component, Input} from '@angular/core';
import {MatCardModule} from '@angular/material/card';
import {Task} from '@client/entities/task';
import {DatePipe} from '@angular/common';

@Component({
    selector: "ts-task-card",
    template: `
        <mat-card appearance="outlined">
            <mat-card-header>
                <mat-card-title>{{ task.title }}</mat-card-title>
            </mat-card-header>
            <mat-card-footer class="p-4 flex items-center gap-2">
                <div class="text-sm">
                    {{ task.created_at | date }} <br>
                    <small class="text-slate-500">{{ task.code }}</small>
                </div>
                <div class="ms-auto">
                    <img src="/assets/img/1.png" class="rounded-full w-6 border border-1 border-slate-400">
                </div>
            </mat-card-footer>
        </mat-card>
    `,
    imports: [MatCardModule, DatePipe]
})
export class TaskCardComponent {
    @Input() task: Task
}
