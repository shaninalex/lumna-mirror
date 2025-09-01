import {Component, inject, OnInit} from '@angular/core';
import {ActivatedRoute} from '@angular/router';
import {Project} from '@client/entities/project';
import {ProjectCardComponent} from '@client/entities/project/ui';

@Component({
    selector: 'fr-projects-list',
    imports: [
        ProjectCardComponent
    ],
    template: `
        <div class="flex flex-col gap-2">
            @for (project of projects; track project.id) {
                <fr-project-card [project]="project" />
            }
        </div>
    `
})
export class ProjectsList implements OnInit {
    activatedRoute = inject(ActivatedRoute)
    projects: Project[] | undefined = this.activatedRoute.snapshot.data['projects']

    ngOnInit() {
        console.log(this.projects)
    }
}
