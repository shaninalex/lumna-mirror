import {Component} from '@angular/core';
import {PageTitleSetter} from '@client/shared/ui';
import {MatCardModule} from '@angular/material/card';
import {MatButtonModule} from '@angular/material/button';
import {RouterLink} from '@angular/router';

@Component({
    selector: "ts-project-list-page",
    template: `
        <div class="grid grid-cols-3 gap-4">
            <mat-card class="example-card" appearance="outlined">
                <mat-card-header>
                    <mat-card-title class="cursor-pointer" [routerLink]="['/projects/123']">Taskiro</mat-card-title>
                    <mat-card-subtitle>Project tracking service</mat-card-subtitle>
                </mat-card-header>
                <mat-card-actions>
                    <a matButton [routerLink]="['/projects/123']">Detail</a>
                </mat-card-actions>
            </mat-card>
            <mat-card class="example-card" appearance="outlined">
                <mat-card-header>
                    <mat-card-title class="cursor-pointer" [routerLink]="['/projects/123']">SoundStream</mat-card-title>
                    <mat-card-subtitle>Music listening application</mat-card-subtitle>
                </mat-card-header>
                <mat-card-actions>
                    <a matButton [routerLink]="['/projects/123']">Detail</a>
                </mat-card-actions>
            </mat-card>
        </div>`,
    imports: [
        MatCardModule,
        MatButtonModule,
        RouterLink,
    ]
})
export class ProjectListPageComponent extends PageTitleSetter {
    pageTitle = "Projects";
}
