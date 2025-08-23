import {Component} from '@angular/core';
import {PageTitleSetter} from '@client/shared/ui';
import {MatListModule} from '@angular/material/list';

@Component({
    selector: "jr-page-home",
    imports: [
        MatListModule
    ],
    template: `
        <mat-list>
            <mat-list-item>Latest activities</mat-list-item>
            <mat-list-item>New tasks assignments</mat-list-item>
            <mat-list-item>Reactions</mat-list-item>
            <mat-list-item>Comments</mat-list-item>
            <mat-list-item>Mentions</mat-list-item>
            <mat-list-item>etc...</mat-list-item>
        </mat-list>
    `
})
export class PageHomeComponent extends PageTitleSetter {
    pageTitle = "Home";
}
