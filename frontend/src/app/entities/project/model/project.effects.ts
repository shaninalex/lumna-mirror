import { Actions, createEffect, ofType } from '@ngrx/effects';
import { inject } from '@angular/core';
import { Router } from '@angular/router';
import { actionProjectList, actionProjectsSet } from './project.actions';
import { exhaustMap, of, switchMap } from 'rxjs';
import { ProjectService } from '../api/project.service';

//  TODO: class based efffect
export const ProjectsGetEffect = createEffect(
    (actions$ = inject(Actions), service = inject(ProjectService), router = inject(Router)) => {
        return actions$.pipe(
            ofType(actionProjectList.type),
            exhaustMap(() =>
                service
                    .GetProjects()
                    .pipe(switchMap((data) => of(actionProjectsSet({ projects: data })))),
            ),
        );
    },
    { functional: true, dispatch: true },
);
