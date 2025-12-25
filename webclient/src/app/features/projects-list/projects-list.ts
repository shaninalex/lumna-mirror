import {Component, inject} from '@angular/core';
import {ProjectStore} from '@entities/project';
import {JsonPipe} from '@angular/common';
import {ProjectCard} from '@entities/project/ui/project-card/project-card';

@Component({
    selector: 'app-projects-list-feature',
    imports: [JsonPipe, ProjectCard],
    templateUrl: './projects-list.html',
    styleUrl: './projects-list.css',
})
export class ProjectsListFeature {
    readonly projectStore = inject(ProjectStore);
    readonly projects = this.projectStore.projects;
}
