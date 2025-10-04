import {Component, Input} from '@angular/core';
import {Project} from '@client/entities/project';
import {RouterLink} from '@angular/router';
import {DatePipe} from '@angular/common';
import {MatCardModule} from '@angular/material/card';

@Component({
    selector: "fr-project-card",
    imports: [
        RouterLink,
        DatePipe,
        MatCardModule
    ],
    template: `
        <mat-card class="example-card" appearance="outlined">
            <mat-card-header>
                <mat-card-title>
                    <a [routerLink]="[ project.code ]">
                        {{ project.title }}
                    </a>
                </mat-card-title>
            </mat-card-header>
            <mat-card-content class="mt-4">
                <div>Last updated: {{ project.updated_at | date }}</div>
                <div class="text-gray-400 text-sm">[{{ project.code }}]</div>
            </mat-card-content>
        </mat-card>

    `
})
export class ProjectCardComponent {
    @Input() project: Project
}
