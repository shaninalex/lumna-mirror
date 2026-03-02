import {Component, inject, OnInit} from '@angular/core';
import {ActivatedRoute, ActivatedRouteSnapshot} from '@angular/router';
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
        <hr />
        <app-project-delete-feature [project$]="project$" />
    `,
})
export class ProjectEditPage {
    private ui = inject(UiService);
    private route = inject(ActivatedRoute);

    project$: Observable<ProjectModel> = this.route.parent!.data.pipe(
        filter((data) => !!data['project']),
        map((data) => data['project'] as ProjectModel),
        tap((project) => this.ui.setPageTitle(`Project: ${project.title}`)),
    );
}
