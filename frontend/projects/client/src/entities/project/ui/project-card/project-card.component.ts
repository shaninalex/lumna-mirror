import {Component, Input} from '@angular/core';
import {Project} from '@client/entities/project';
import {RouterLink} from '@angular/router';
import {DatePipe} from '@angular/common';

@Component({
    selector: "fr-project-card",
    imports: [
        RouterLink,
        DatePipe
    ],
    template: `
        <a [routerLink]="[ project.project_key ]" class="card bg-base-100 border border-base-300">
            <div class="card-body">
                <div class="font-bold font-lg mb-2">{{ project.title }}</div>
                <div>Last updated: {{ project.updated_at | date }}</div>
                <div class="text-gray-400 text-sm">[{{ project.project_key }}]</div>
            </div>
        </a>
    `
})
export class ProjectCardComponent {
    @Input() project: Project
}
