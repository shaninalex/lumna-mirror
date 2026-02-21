import { Component, inject } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { ProjectModel } from '@entities/project';
import { ProjectDeleteFeature, ProjectEditFeature } from '@features';
import { UiService } from '@shared/ui';
import { filter, map, Observable, tap } from 'rxjs';

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
export class ProjectEditPage {
    private ui = inject(UiService);

    private route = inject(ActivatedRoute);
    project$: Observable<ProjectModel> = this.route.data.pipe(
        filter((data) => !!data['project']),
        map((data) => data['project'] as ProjectModel),
        tap((project) => this.ui.setPageTitle(`Edit project: ${project.title}`)),
    );
}
