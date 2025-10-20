import { Component, inject } from '@angular/core';
import { Project } from '@client/entities/project';
import { ProjectCardComponent } from '@client/entities/project/ui';
import { Store } from '@ngrx/store';
import { AppState } from '@client/shared/store';
import { Observable } from 'rxjs';
import { selectProjects } from '@client/entities/project/model/project.selectors';
import { AsyncPipe } from '@angular/common';
import { NewProjectFormComponent } from '@client/features/project/project-list-feature/new-project-form.component';

@Component({
    selector: 'lu-project-list-feature',
    imports: [
        ProjectCardComponent,
        AsyncPipe,
        NewProjectFormComponent,
    ],
    template: `
        <div class="mb-4">
            <lu-new-project-form />
        </div>

        @if (projects$ | async; as projects) {
            <div class="flex flex-col gap-2 overflow-auto">
                @if (!projects?.length) {
                    there are no projects yet
                }
                @for (project of projects; track project.id) {
                    <lu-project-card [project]="project"/>
                }
            </div>
        }
    `,
})
export class ProjectListFeatureComponent {
    private store = inject(Store<AppState>);
    projects$: Observable<Project[]> = this.store.select(selectProjects);
}
