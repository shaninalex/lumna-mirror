import { Component, inject } from '@angular/core';
import { Store } from '@ngrx/store';
import { Dialog } from '@angular/cdk/dialog';

import { ProjectPayload, ProjectState, actionProjectCreate } from '@entities/project';
import { ProjectForm } from './components';

@Component({
    selector: 'app-projects-create-feature',
    template: `
        <div>
            <button
                (click)="newProject()"
                class="bg-teal-500 text-white rounded-lg px-4 py-2 cursor-pointer"
            >
                New project
            </button>
        </div>
    `,
})
export class ProjectCreateFeature {
    dialog = inject(Dialog);

    private store = inject(Store<ProjectState>);

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
