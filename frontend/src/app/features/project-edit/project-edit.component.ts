import { Component, inject, Input, OnInit, signal } from '@angular/core';
import { Store } from '@ngrx/store';
import { Actions, ofType } from '@ngrx/effects';
import { form, required, FormField } from '@angular/forms/signals';

import {
    actionProjectUpdate,
    actionProjectUpsert,
    ProjectModel,
    ProjectPayload,
    ProjectState,
} from '@entities/project';
import { filter, Observable, tap } from 'rxjs';
import { AsyncPipe } from '@angular/common';

@Component({
    selector: 'app-project-edit-feature',
    imports: [FormField, AsyncPipe],
    templateUrl: './project-edit.component.html',
})
export class ProjectEditFeature implements OnInit {
    private actions$ = inject(Actions);
    private store = inject(Store<ProjectState>);
    private projectId: string;
    @Input() project$: Observable<ProjectModel>;

    loading = false;

    projectFormModel = signal<ProjectPayload>({
        title: '',
    });
    errors = signal<string[]>([]);

    ngOnInit() {
        this.project$.subscribe((project) => {
            this.projectId = project.id;
            this.projectFormModel.set({ title: project.title });
        });
        this.actions$.pipe(ofType(actionProjectUpsert)).subscribe(() => (this.loading = false));
    }

    projectForm = form(this.projectFormModel, (schemaPath) => {
        required(schemaPath.title, { message: 'Title is required' });
    });

    submit(event: Event): void {
        event.preventDefault();
        this.loading = true;
        this.store.dispatch(
            actionProjectUpdate({ id: this.projectId, data: this.projectFormModel() }),
        );
    }
}
