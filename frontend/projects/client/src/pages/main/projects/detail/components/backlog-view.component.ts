import {Component} from '@angular/core';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatInput, MatInputModule} from '@angular/material/input';
import {MatIconModule} from '@angular/material/icon';
import {MatButtonModule} from '@angular/material/button';
import {SprintEditListCardComponent} from '@client/features/sprint';

@Component({
    selector: "ts-backlog-view",
    imports: [
        MatFormFieldModule,
        MatInputModule,
        MatIconModule,
        MatButtonModule,
        SprintEditListCardComponent,
    ],
    template: `
        <div class="flex items-center gap-4 mb-4">
            <mat-form-field appearance="outline">
                <mat-label>
                    Search Backlog
                </mat-label>
                <button mat-icon-button matPrefix class="ml-1">
                    <mat-icon>search</mat-icon>
                </button>
                <input matInput type="search">
            </mat-form-field>

            <div class="flex items-center">
                <img src="assets/img/1.png" alt="" class="rounded-full w-8 border border-2 border-slate-400">
                <svg viewBox="-4 -4 24 24" class="-ml-2 rounded-full w-8 border border-2 border-slate-400 bg-slate-200"
                     role="presentation">
                    <path fill="currentcolor" fill-rule="evenodd"
                          d="M8 1.5a2.5 2.5 0 1 0 0 5 2.5 2.5 0 0 0 0-5M4 4a4 4 0 1 1 8 0 4 4 0 0 1-8 0m-2 9a3.75 3.75 0 0 1 3.75-3.75h4.5A3.75 3.75 0 0 1 14 13v2h-1.5v-2a2.25 2.25 0 0 0-2.25-2.25h-4.5A2.25 2.25 0 0 0 3.5 13v2H2z"
                          clip-rule="evenodd"></path>
                </svg>
            </div>
        </div>

        <div class="flex flex-col gap-4">
            <ts-sprint-edit-list-card/>
            <ts-sprint-edit-list-card/>
        </div>
    `
})
export class BacklogViewComponent {

}
