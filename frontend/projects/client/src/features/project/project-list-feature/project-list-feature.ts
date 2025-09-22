import {Component, inject} from '@angular/core';
import {CreateProjectAction, Project, SetProjectAction} from '@client/entities/project';
import {ProjectCardComponent} from '@client/entities/project/ui';
import {Store} from '@ngrx/store';
import {AppState} from '@client/shared/store';
import {Observable} from 'rxjs';
import {selectProjects} from '@client/entities/project/model/project.selectors';
import {AsyncPipe} from '@angular/common';
import {NewProjectFormComponent} from '@client/features/project/project-list-feature/new-project-form.component';
import {Actions, ofType} from '@ngrx/effects';

@Component({
    selector: 'fr-project-list-feature',
    imports: [
        ProjectCardComponent,
        AsyncPipe,
        NewProjectFormComponent,
    ],
    template: `
        <div class="mb-4">
            @if (projectForm) {
                <fr-new-project-form [loading]="loading" (onCancel)="toggleProjectForm()"
                                     (onSubmit)="onSubmit($event)"/>
            } @else {
                <button class="btn btn-primary" (click)="toggleProjectForm()">Create</button>
            }
        </div>

        @if (projects$ | async; as projects) {
            <div class="flex flex-col gap-2 overflow-auto">
                @if (!projects?.length) {
                    there are no projects yet
                }
                @for (project of projects; track project.id) {
                    <fr-project-card [project]="project"/>
                }
            </div>
        }
    `,
})
export class ProjectListFeatureComponent {
    private store = inject(Store<AppState>);
    private actions$ = inject(Actions);
    projects$: Observable<Project[]> = this.store.select(selectProjects);
    projectForm: boolean = false;
    loading: boolean = false;

    constructor() {
        this.actions$.pipe(ofType(SetProjectAction)).subscribe(() => {
            this.loading = this.projectForm = false
        })
    }

    toggleProjectForm(): void {
        this.projectForm = !this.projectForm;
    }

    onSubmit(title: string): void {
        this.loading = true;
        const project: Record<string, string> = { title }
        this.store.dispatch(CreateProjectAction({payload: project}))
    }
}
