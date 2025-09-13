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
        <a [routerLink]="[ project.project_key ]" class="block border-lg bg-white p-4 border rounded-lg border-slate-300 cursor-pointer">
            <div class="font-bold font-lg mb-2">{{ project.title }}</div>
            <div>Last updated: {{ project.updated_at | date }}</div>
            <div class="text-slate-400 text-sm">[{{ project.project_key }}]</div>
        </a>
    `
})
export class ProjectCardComponent {
    @Input() project: Project
}
