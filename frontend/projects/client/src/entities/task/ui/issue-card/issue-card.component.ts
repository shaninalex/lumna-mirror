import {Component, Input, OnInit} from '@angular/core';
import {MatCardModule} from '@angular/material/card';
import {Task} from '@client/entities/task';
import {DatePipe} from '@angular/common';

@Component({
    selector: "ts-issue-card",
    template: `
        <mat-card appearance="outlined">
            <mat-card-header>
                <h4 class="text-sm">{{ task.title }}</h4>
            </mat-card-header>
            <mat-card-footer class="p-4 flex items-center gap-2">
                <div class="text-sm">{{ task.created_at | date }}</div>
                <div class="ms-auto">
                    <img src="/assets/img/1.png" class="rounded-full w-6 border border-1 border-slate-400">
                </div>
            </mat-card-footer>
        </mat-card>
    `,
    imports: [MatCardModule, DatePipe]
})
export class IssueCardComponent implements OnInit {
    @Input() task: Task

    ngOnInit() {
        console.log(this.task.created_at)
    }
}
