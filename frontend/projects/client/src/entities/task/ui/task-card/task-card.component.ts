import {Component, inject, Input} from '@angular/core';
import {MatCardModule} from '@angular/material/card';
import {Task} from '@client/entities/task';
import {DatePipe} from '@angular/common';
import {MatButtonModule} from '@angular/material/button';
import {MatDialog, MatDialogModule} from '@angular/material/dialog';
import {TaskDetailViewComponent} from './task-detail-view.component';
import {ScrollStrategy} from '@angular/cdk/overlay';

@Component({
    selector: "ts-task-card",
    template: `
        <mat-card appearance="outlined">
            <mat-card-header>
                <button class="cursor-pointer" (click)="openDetailView()" type="button">
                    {{ task.title }}
                </button>
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
    imports: [MatCardModule, DatePipe, MatButtonModule, MatDialogModule]
})
export class TaskCardComponent {
    @Input() task: Task;
    readonly dialog = inject(MatDialog);

    openDetailView() {
        this.dialog.open(TaskDetailViewComponent, {
            data: {task: this.task},
            maxWidth: "100%",
            panelClass: "lg-dialog"
        });
    }
}
