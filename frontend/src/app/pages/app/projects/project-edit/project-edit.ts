import { Component, inject } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { ProjectEditFeature, ProjectDeleteFeature } from '@features/index';

@Component({
    selector: 'app-project-edit',
    imports: [ProjectEditFeature, ProjectDeleteFeature],
    template: `
        <h1>Project Edit</h1>
        <app-project-edit-feature [projectId]="projectId" />
        <div class="mb-4"></div>
        <app-project-delete-feature [projectId]="projectId" />
    `,
})
export class ProjectEdit {
    projectId: string;
    private activatedRoute = inject(ActivatedRoute);

    constructor() {
        // Access route parameters
        this.activatedRoute.params.subscribe((params) => {
            try {
                const id = params['id'];
                this.projectId = id;
            } catch {
                alert('Invalid project id');
            }
        });
    }
}
