import { Component, inject, signal } from '@angular/core';
import { ProjectCreateFeature } from '@features';
import { UiService } from '@shared/ui';
import { Store } from '@ngrx/store';
import { ProjectCard, ProjectModel, ProjectState, selectProjects } from '@entities/project';
import { Observable } from 'rxjs';
import { AsyncPipe } from '@angular/common';
import {projectListGrid} from './grid';


@Component({
    selector: 'app-project-list',
    imports: [AsyncPipe, ProjectCard, ProjectCreateFeature, projectListGrid],
    template: `<div class="is-flex is-flex-direction-column is-gap-2">
        <div class="is-flex is-align-items-center is-justify-content-space-between">
            <app-projects-create-feature />
            <button (click)="toggleViewMode()" class="button is-size-6">
                @if (viewMode() === 'grid') {
                    <i class="fa-solid fa-grip"></i>
                } @else {
                    <i class="fa-solid fa-list"></i>
                }
            </button>
        </div>

        @if (projects | async; as projects) {
            @if (viewMode() == 'grid') {
                <app-project-list--grid>
                    @for (p of projects; track p.id) {
                        <app-project-card [project]="p" />
                    }
                </app-project-list--grid>
            } @else {
                <div class="is-flex is-flex-direction-column is-gap-2">
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
