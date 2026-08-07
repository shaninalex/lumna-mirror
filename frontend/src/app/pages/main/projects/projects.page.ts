import { Component } from '@angular/core';
import { GlobalLayout } from '@core/layout';
import { RouterLink } from "@angular/router";

@Component({
    selector: 'lu-projects-page',
    imports: [GlobalLayout, RouterLink],
    template: `
        <lu-global-layout>
            <div class="container py-4">
                <div class="d-flex justify-content-between align-items-center mb-4">
                    <div>
                        <div class="text-muted small">Workspace</div>
                        <div class="d-flex align-items-center gap-3">
                            <h2 class="mb-0">Lumna Dev</h2>

                            <a routerLink="/workspaces" class="text-decoration-none small">
                                Switch workspace
                            </a>
                        </div>
                    </div>

                    <button class="btn btn-primary">
                        New Project
                    </button>
                </div>

                <div class="row g-4">
                    <div class="col-12 col-md-6 col-xl-4">
                        <div class="card h-100">
                            <div class="card-body">

                                <div class="d-flex justify-content-between align-items-start mb-3">
                                    <div>
                                        <h5 class="card-title mb-1">
                                            Website Redesign
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

                                <div class="d-flex justify-content-between text-muted small mb-3">
                                    <span>12 members</span>
                                    <span>245 tasks</span>
                                </div>

                            </div>

                            <div class="card-footer bg-transparent border-0 pt-0">
                                <button class="btn btn-outline-primary w-100">
                                    Open project
                                </button>
                            </div>
                        </div>
                    </div>

                    <div class="col-12 col-md-6 col-xl-4">
                        <div class="card h-100">
                            <div class="card-body">

                                <div class="d-flex justify-content-between align-items-start mb-3">
                                    <div>
                                        <h5 class="card-title mb-1">
                                            Mobile App
                                        </h5>

                                        <div class="text-muted small">
                                            MOBILE
                                        </div>
                                    </div>

                                    <span class="badge text-bg-warning">
                                        Planning
                                    </span>
                                </div>

                                <p class="card-text text-muted">
                                    Placeholder description for the project.
                                </p>

                                <div class="d-flex justify-content-between text-muted small mb-3">
                                    <span>6 members</span>
                                    <span>38 tasks</span>
                                </div>

                            </div>

                            <div class="card-footer bg-transparent border-0 pt-0">
                                <button class="btn btn-outline-primary w-100">
                                    Open project
                                </button>
                            </div>
                        </div>
                    </div>

                    <div class="col-12 col-md-6 col-xl-4">
                        <div class="card h-100">
                            <div class="card-body">

                                <div class="d-flex justify-content-between align-items-start mb-3">
                                    <div>
                                        <h5 class="card-title mb-1">
                                            Internal Tools
                                        </h5>

                                        <div class="text-muted small">
                                            PLATFORM
                                        </div>
                                    </div>

                                    <span class="badge text-bg-secondary">
                                        Archived
                                    </span>
                                </div>

                                <p class="card-text text-muted">
                                    Placeholder description for the project.
                                </p>

                                <div class="d-flex justify-content-between text-muted small mb-3">
                                    <span>4 members</span>
                                    <span>91 tasks</span>
                                </div>

                            </div>

                            <div class="card-footer bg-transparent border-0 pt-0">
                                <button class="btn btn-outline-primary w-100">
                                    Open project
                                </button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>  
        </lu-global-layout>
    `,
})
export class ProjectsPage {

}
