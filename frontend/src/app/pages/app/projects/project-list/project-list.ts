import { Component, inject, signal } from '@angular/core';
import { ProjectCreateFeature } from '@features';
import { UiService } from '@shared/ui';
import { Store } from '@ngrx/store';
import { ProjectCard, ProjectModel, ProjectState, selectProjects } from '@entities/project';
import { Observable } from 'rxjs';
import { AsyncPipe } from '@angular/common';


@Component({
    selector: 'app-project-list',
    imports: [AsyncPipe, ProjectCard, ProjectCreateFeature],
    template: `<div class="d-flex flex-column gap-4">
        <div class="d-flex align-items-center justify-content-between">
            <app-projects-create-feature />
            <button (click)="toggleViewMode()" class="btn btn-outline-secondary btn-sm">
                @if (viewMode() === 'grid') {
                    <i class="fa-solid fa-grip"></i>
                } @else {
                    <i class="fa-solid fa-list"></i>
                }
            </button>
        </div>

        @if (projects | async; as projects) {
            @if (viewMode() == 'grid') {
                <div class="row">
                    @for (p of projects; track p.id) {
                        <div class="col-md-4 mb-4">
                            <app-project-card [project]="p" />
                        </div>
                    }
                </div>
            } @else {
                <div class="d-flex flex-column gap-4">
                    @for (p of projects; track p.id) {
                        <app-project-card [project]="p" />
                    }
                </div>
            }
        }
    </div>`,
})
export class ProjectList {
    private ui = inject(UiService);

    private store = inject(Store<ProjectState>);
    projects: Observable<ProjectModel[]> = this.store.select(selectProjects);

    viewMode = signal<'list' | 'grid'>('grid');

    constructor() {
        this.ui.setPageTitle('Projects');
    }

    toggleViewMode(): void {
        if (this.viewMode() === 'list') {
            this.viewMode.set('grid');
        } else {
            this.viewMode.set('list');
        }
    }
}
