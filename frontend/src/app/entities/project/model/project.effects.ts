import { Actions, createEffect, ofType } from '@ngrx/effects';
import { inject, Injectable } from '@angular/core';
import { actionProjectList, actionProjectsSet } from './project.actions';
import { exhaustMap, of, switchMap } from 'rxjs';
import { ProjectService } from '../api/project.service';

@Injectable()
export class ProjectsEffects {
    private actions$ = inject(Actions);
    private projectsApi = inject(ProjectService);

    get_projects_list$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionProjectList.type),
            exhaustMap(() =>
                this.projectsApi
                    .GetProjects()
                    .pipe(switchMap((data) => of(actionProjectsSet({ projects: data })))),
            ),
        ),
    );
}
