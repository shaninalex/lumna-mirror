import { Component, inject, signal } from '@angular/core';
import { ProjectCreateFeature } from '@features/index';
import { UiService } from '@shared/ui';
import { Store } from '@ngrx/store';
import { ProjectCard, ProjectModel, ProjectState, selectProjects } from '@entities/project';
import { Observable } from 'rxjs';
import { AsyncPipe, NgClass } from '@angular/common';

@Component({
    selector: 'app-project-list',
    imports: [NgClass, AsyncPipe, ProjectCard, ProjectCreateFeature],
    template: `<div class="is-flexflex-col gap-4">
        <app-projects-create-feature />

        @if (projects | async; as projects) {
            <div
                [ngClass]="{
                    'grid grid-cols-3 gap-6': viewMode() === 'grid',
                    'is-flexflex-col gap-4': viewMode() === 'list',
                }"
            >
                @for (p of projects; track p.id) {
                    <app-project-card [project]="p" />
                }
            </div>
        }

        <div>
            <button (click)="toggleViewMode()" class="button is-small btn-secondary">
                @if (viewMode() === 'grid') {
                    <i class="fa-solid fa-grip"></i>
                } @else {
                    <i class="fa-solid fa-list"></i>
                }
            </button>
        </div>
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
