import { Actions, createEffect, ofType } from '@ngrx/effects';
import { inject, Injectable } from '@angular/core';
import {
    actionProjectCreate,
    actionProjectList,
    actionProjectsAdd,
    actionProjectsSetList,
} from './project.actions';
import { exhaustMap, map, of, switchMap, tap } from 'rxjs';
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
                    .pipe(switchMap((data) => of(actionProjectsSetList({ projects: data })))),
            ),
        ),
    );

    create_project$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionProjectCreate.type),
            exhaustMap((action) =>
                this.projectsApi
                    .CreateProject(action.payload)
                    .pipe(switchMap((data) => of(actionProjectsAdd({ project: data })))),
            ),
        ),
    );
}
