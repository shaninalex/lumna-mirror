import {Component, Input} from '@angular/core';
import {Project} from '@client/entities/project';
import {RouterLink} from '@angular/router';
import {DatePipe} from '@angular/common';

@Component({
    selector: "lu-project-card",
    imports: [
        RouterLink,
        DatePipe,
    ],
    template: `
        <div class="card">
            <div class="card-title">
                <a [routerLink]="[ project.code ]">
                    {{ project.title }}
                </a>
            </div>
            <div class="mt-4">
                <div>Last updated: {{ project.updated_at | date }}</div>
                <div class="text-gray-400 text-sm">[{{ project.code }}]</div>
            </div>
        </div>

    `
})
export class ProjectCardComponent {
    @Input() project: Project
}
