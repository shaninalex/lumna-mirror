import {Component, Input} from '@angular/core';
import {ProjectModel} from '@entities/project';
import {RouterLink} from '@angular/router';

@Component({
    selector: 'app-project-card',
    imports: [
        RouterLink
    ],
    template: `
        <a [routerLink]="['/projects', project.id]"
           class="bg-gray-100 rounded-xl p-4 block">
            <div class="flex justify-between">
                <div class="flex justify-between items-center">
                    <i class="fa-regular fa-calendar-days"></i>
                    <div>{{ project.name }}</div>
                </div>
                <div>
                    <i class="fa-solid fa-ellipsis"></i>
                </div>
            </div>
        </a>
    `,
    styleUrl: './project-card.css',
})
export class ProjectCard {
    @Input() project: ProjectModel
}
