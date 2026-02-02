import { Component, inject, signal } from '@angular/core';
import { Observable } from 'rxjs';
import { AsyncPipe, NgClass } from '@angular/common';
import { Dialog } from '@angular/cdk/dialog';
import { Store } from '@ngrx/store';

import { ProjectCard } from '@entities/project/ui';
import { ProjectForm } from './components';
import {
    ProjectModel,
    ProjectPayload,
    ProjectState,
    selectProjects,
    actionProjectCreate,
} from '@entities/project';

@Component({
    selector: 'app-projects-list-feature',
    imports: [ProjectCard, NgClass, AsyncPipe],
    templateUrl: './projects-list.html',
})
export class ProjectsListFeature {
    private store = inject(Store<ProjectState>);
    projects: Observable<ProjectModel[]> = this.store.select(selectProjects);

    viewMode = signal<'list' | 'grid'>('grid');
    dialog = inject(Dialog);

    toggleViewMode(): void {
        if (this.viewMode() === 'list') {
            this.viewMode.set('grid');
        } else {
            this.viewMode.set('list');
        }
    }

    newProject(): void {
        const dialogRef = this.dialog.open<ProjectPayload>(ProjectForm, {
            width: '250px',
        });

        dialogRef.closed.subscribe((result) => {
            if (!result) return;
            const payload: ProjectPayload = {
                title: result.title,
            };
            this.store.dispatch(actionProjectCreate({ payload }));
        });
    }
}
