import {Component, Input} from '@angular/core';
import {MatCardModule} from '@angular/material/card';
import {Issue} from '@client/entities/issue';

@Component({
    selector: "ts-issue-card",
    template: `
        <mat-card appearance="outlined">
            <mat-card-header>
                <h4 class="text-base">{{ task.title }}</h4>
            </mat-card-header>
            <mat-card-footer class="p-4 flex items-center gap-2">
<!--                <mat-chip-set>-->
<!--                    <mat-chip>tags</mat-chip>-->
<!--                </mat-chip-set>-->
<!--                <mat-icon>flag</mat-icon> &lt;!&ndash; priority &ndash;&gt;-->
                <div class="ms-auto">
                    <img src="/assets/img/1.png" class="rounded-full w-6 border border-1 border-slate-400">
                </div>
            </mat-card-footer>
        </mat-card>
    `,
    imports: [MatCardModule]
})
export class IssueCardComponent {
    @Input() task: Issue
}
