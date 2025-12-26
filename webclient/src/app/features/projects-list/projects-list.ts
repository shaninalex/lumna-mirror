import {Component, inject, signal} from '@angular/core';
import {projectEvents, ProjectPayload, ProjectStore} from '@entities/project';
import {ProjectCard} from '@entities/project/ui';
import {NgClass} from '@angular/common';
import {ProjectForm} from '@features/project-form/project-form';
import {Dialog} from '@angular/cdk/dialog';
import {Dispatcher} from '@ngrx/signals/events';

@Component({
    selector: 'app-projects-list-feature',
    imports: [ProjectCard, NgClass],
    templateUrl: './projects-list.html',
    styleUrl: './projects-list.css',
})
export class ProjectsListFeature {
    private readonly dispatcher = inject(Dispatcher)
    private readonly projectStore = inject(ProjectStore);
    readonly projects = this.projectStore.entities;

    viewMode = signal<"list" | "grid">("grid");
    dialog = inject(Dialog);

    toggleViewMode(): void {
        if (this.viewMode() === "list") {
            this.viewMode.set("grid")
        } else {
            this.viewMode.set("list")
        }
    }

    newProject(): void {
        const dialogRef = this.dialog.open<ProjectPayload>(ProjectForm, {
            width: '250px',
        });

        dialogRef.closed.subscribe(result => {
            if (result) this.dispatcher.dispatch(projectEvents.createProject(result))
        });
    }

    handleOnDeleteProject(projectId: number): void {
        this.dispatcher.dispatch(projectEvents.deleteProject(projectId))
    }
}
