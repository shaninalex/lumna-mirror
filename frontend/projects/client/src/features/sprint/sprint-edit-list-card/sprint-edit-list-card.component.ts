import {Component} from '@angular/core';
import {MatCardModule} from '@angular/material/card';
import {MatChipsModule} from '@angular/material/chips';
import {MatButtonModule} from '@angular/material/button';
import {IssuesEditTableComponent} from '@client/features/issue';

@Component({
    selector: "ts-sprint-edit-list-card",
    template: `
        <mat-card class="example-card" appearance="outlined">
            <mat-card-header>
                <mat-card-title>Setup basic layout</mat-card-title>
            </mat-card-header>
            <mat-card-content>
                <ts-issues-edit-table />
            </mat-card-content>
            <mat-card-footer class="p-4">
                <button matButton="outlined" type="button">create</button>
            </mat-card-footer>
        </mat-card>
    `,
    imports: [
        MatCardModule,
        MatChipsModule,
        MatButtonModule,
        IssuesEditTableComponent,
    ]
})
export class SprintEditListCardComponent {


}
