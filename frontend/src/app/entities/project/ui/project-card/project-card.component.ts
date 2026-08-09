import { Component, Input } from '@angular/core';
import { RouterLink } from '@angular/router';
import { ProjectModel } from '@entities/project/model';

@Component({
    selector: 'lu-project-card',
    imports: [RouterLink],
    template: `
        <div class="card h-100">
            <div class="card-body">

                <div class="d-flex justify-content-between align-items-start mb-3">
                    <div>
                        <h5 class="card-title mb-1">
                            <a routerLink="/">{{ project.title }} </a>
                        </h5>

                        <div class="text-muted small">
                            WEB
                        </div>
                    </div>

                    <span class="badge text-bg-success">
                        Active
                    </span>
                </div>

                <p class="card-text text-muted">
                    Placeholder description for the project. Brief summary of
                    what this project is about.
                </p>

                <div class="d-flex justify-content-between text-muted small">
                    <span>12 members</span>
                    <span>245 tasks</span>
                </div>
            </div>
        </div>
    `,
})
export class ProjectCardComponent {
    @Input() project: ProjectModel;
}
