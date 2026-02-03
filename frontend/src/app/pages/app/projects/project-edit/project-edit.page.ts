import { Component, inject, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { ProjectState, ProjectModel, selectProjectByID } from '@entities/project';
import { ProjectEditFeature, ProjectDeleteFeature } from '@features/index';
import { Store } from '@ngrx/store';
import { UiService } from '@shared/ui';
import { Observable, filter, switchMap, tap } from 'rxjs';

@Component({
    selector: 'app-project-edit',
    imports: [ProjectEditFeature, ProjectDeleteFeature],
    template: `
        <h1>Project Edit</h1>
        <app-project-edit-feature [project$]="project$" />
        <div class="border-b border-slate-200 dark:border-slate-700 my-4"></div>
        <app-project-delete-feature [project$]="project$" />
    `,
})
export class ProjectEditPage implements OnInit {
    private activatedRoute = inject(ActivatedRoute);
    private store = inject(Store<ProjectState>);
    private ui = inject(UiService);
    project$: Observable<ProjectModel>;

    ngOnInit() {
        this.project$ = this.activatedRoute.params.pipe(
            switchMap((params) =>
                this.store.select(selectProjectByID(params['id'])).pipe(
                    filter((project) => !!project),
                    tap((project) => this.ui.setPageTitle(`Edit project: ${project.title}`)),
                ),
            ),
        );
    }
}
