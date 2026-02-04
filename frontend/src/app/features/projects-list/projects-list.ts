import { Component, inject, signal } from '@angular/core';
import { Observable } from 'rxjs';
import { AsyncPipe, NgClass } from '@angular/common';
import { Store } from '@ngrx/store';

import { ProjectCard } from '@entities/project/ui';
import { ProjectModel, ProjectState, selectProjects } from '@entities/project';
import { ProjectCreateFeature } from '@features';

@Component({
    selector: 'app-projects-list-feature',
    imports: [ProjectCard, NgClass, AsyncPipe, ProjectCreateFeature],
    template: `<div class="flex flex-col gap-4">
        <app-projects-create-feature />

        @if (projects | async; as projects) {
            <div
                [ngClass]="{
                    'grid grid-cols-3 gap-6': viewMode() === 'grid',
                    'flex flex-col gap-4': viewMode() === 'list',
                }"
            >
                @for (p of projects; track p.id) {
                    <app-project-card [project]="p" />
                }
            </div>
        }

        <div>
            <button (click)="toggleViewMode()" class="btn btn-sm btn-secondary">
                @if (viewMode() === 'grid') {
                    <i class="fa-solid fa-grip"></i>
                } @else {
                    <i class="fa-solid fa-list"></i>
                }
            </button>
        </div>
    </div>`,
})
export class ProjectsListFeature {
    private store = inject(Store<ProjectState>);
    projects: Observable<ProjectModel[]> = this.store.select(selectProjects);

    viewMode = signal<'list' | 'grid'>('grid');

    toggleViewMode(): void {
        if (this.viewMode() === 'list') {
            this.viewMode.set('grid');
        } else {
            this.viewMode.set('list');
        }
    }
}
