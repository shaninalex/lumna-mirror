import {ChangeDetectionStrategy, Component, inject, Input} from '@angular/core';
import {MatCardModule} from '@angular/material/card';
import {Task} from '@client/entities/task';
import {DatePipe} from '@angular/common';
import {MatButtonModule} from '@angular/material/button';
import {RouterLink} from '@angular/router';

@Component({
    selector: "ts-task-card",
    template: `
        <mat-card appearance="outlined">
            <mat-card-header>
                <a [routerLink]="['/projects', projectCode, task.code]" class="cursor-pointer text-left">
                    {{ task.title }}
                </a>
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
    imports: [MatCardModule, DatePipe, MatButtonModule, RouterLink],
})
export class TaskCardComponent {
    @Input() task: Task;
    @Input() projectCode: string;
}
