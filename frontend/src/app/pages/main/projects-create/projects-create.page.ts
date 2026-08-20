import { Component } from '@angular/core';
import { GlobalLayout } from '@core/layout';
import { ProjectCreateFeature } from '@features/project';

@Component({
    selector: 'lu-projects-create-page',
    imports: [GlobalLayout, ProjectCreateFeature],
    template: `
        <lu-global-layout>
            <div class="container py-4">
                <div class="d-flex justify-content-between align-items-center mb-4">
                    <h2 class="mb-1">Create Project</h2>
                </div>

                <lu-project-create-feature />
            </div>
        </lu-global-layout>
    `,
})
export class ProjectsCreatePage {}
