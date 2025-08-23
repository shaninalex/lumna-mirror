import {Component, inject} from '@angular/core';
import {Project, ProjectService} from '@client/entities/project';
import {MatCardModule} from '@angular/material/card';
import {MatButtonModule} from '@angular/material/button';
import {RouterLink} from '@angular/router';
import {Observable} from 'rxjs';
import {AsyncPipe} from '@angular/common';

@Component({
    selector: "ts-projects-list",
    template: `
        <div class="grid grid-cols-3 gap-4">
            @if (projects | async; as projects) {
                @for (project of projects; track project.id) {
                    <mat-card class="example-card" appearance="outlined">
                        <mat-card-header>
                            <mat-card-title class="cursor-pointer" [routerLink]="['/projects', project.project_key]">
                                {{ project.title }}
                            </mat-card-title>
                            <mat-card-subtitle>Project tracking service</mat-card-subtitle>
                        </mat-card-header>
                        <mat-card-actions>
                            <a matButton [routerLink]="['/projects', project.project_key]">Detail</a>
                        </mat-card-actions>
                    </mat-card>
                }
            }
        </div>`,
    imports: [
        MatCardModule,
        MatButtonModule,
        RouterLink,
        AsyncPipe,
    ]
})
export class ProjectListComponent {
    api = inject(ProjectService)

    // TODO: move to NGRX
    projects: Observable<Array<Project>> = this.api.List();
}
