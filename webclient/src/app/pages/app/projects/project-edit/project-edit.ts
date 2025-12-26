import { Component, inject } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { ProjectStore } from '@entities/project';
import { ProjectEditFeature, ProjectDeleteFeature } from '@features/index';

@Component({
    selector: 'app-project-edit',
    imports: [
        ProjectEditFeature,
        ProjectDeleteFeature
    ],
    template: `
        <h1>Project Edit</h1>
        <app-project-edit-feature [projectId]="projectId" />
        <div class="mb-4"></div>
        <app-project-delete-feature [projectId]="projectId" />
    `,
})
export class ProjectEdit {
    projectId: number
    private activatedRoute = inject(ActivatedRoute);
    private readonly projectStore = inject(ProjectStore);

    constructor() {
        // Access route parameters
        this.activatedRoute.params.subscribe((params) => {
            try {
                const id = parseInt(params['id'])
                this.projectId = id;
            } catch {
                alert("Invalid project id")
            }
        });
    }
}
