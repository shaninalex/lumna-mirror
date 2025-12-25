import {Component, inject, signal} from '@angular/core';
import {ProjectStore} from '@entities/project';
import {ProjectCard} from '@entities/project/ui/project-card/project-card';
import {NgClass} from '@angular/common';

@Component({
    selector: 'app-projects-list-feature',
    imports: [ProjectCard, NgClass],
    templateUrl: './projects-list.html',
    styleUrl: './projects-list.css',
})
export class ProjectsListFeature {
    readonly projectStore = inject(ProjectStore);
    readonly projects = this.projectStore.projects;

    viewMode = signal<"list" | "grid">("grid");

    toggleViewMode(): void {
        if (this.viewMode() === "list") {
            this.viewMode.set("grid")
        } else {
            this.viewMode.set("list")
        }
    }
}
