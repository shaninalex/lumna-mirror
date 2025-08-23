import {Component, Input} from '@angular/core';
import {MatCardModule} from '@angular/material/card';
import {MatChipsModule} from '@angular/material/chips';
import {MatIconModule} from '@angular/material/icon';

@Component({
    selector: "ts-issue-card",
    template: `
        <mat-card appearance="outlined">
            <mat-card-header>
                <h4 class="text-base">[#{{ id }}] Create basic layout structure</h4>
            </mat-card-header>
            <!-- <mat-card-content></mat-card-content> -->
            <mat-card-footer class="p-4 flex items-center gap-2">
                <mat-chip-set>
                    <mat-chip>tags</mat-chip>
                </mat-chip-set>
                <mat-icon>flag</mat-icon> <!-- priority -->
                <div class="ms-auto">
                    <img src="/assets/img/1.png" class="rounded-full w-6 border border-1 border-slate-400">
                </div>
            </mat-card-footer>
        </mat-card>
    `,
    imports: [MatCardModule, MatChipsModule, MatIconModule]
})
export class IssueCardComponent {
    @Input() id: string
}
