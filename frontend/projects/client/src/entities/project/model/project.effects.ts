import {Actions, createEffect, ofType} from '@ngrx/effects';
import {inject} from '@angular/core';
import {HttpErrorResponse} from '@angular/common/http';
import {Router} from '@angular/router';
import {
    CreateProjectAction,
    GetProjectsAction,
    SetProjectAction,
    SetProjectsAction
} from '@client/entities/project/model/project.actions';
import {catchError, EMPTY, exhaustMap, map, of} from 'rxjs';
import {ProjectService} from '@client/entities/project/api/project.service';

export const getProjectsEffect = createEffect(
    (
        actions$ = inject(Actions),
        service = inject(ProjectService),
        router = inject(Router),
    ) => {
        return actions$.pipe(
            ofType(GetProjectsAction.type),
            exhaustMap(() => service.List().pipe(
                map(data => SetProjectsAction({payload: data})),

                // This errors we can catch in global.interceptor
                catchError((error: HttpErrorResponse) => {
                    if (error.status === 401) {
                        router.navigate(['/auth/login']);
                        return of({type: "[session] error"});
                    }
                    return of({type: "[session] error"});
                }),
            ))
        );
    },
    {functional: true, dispatch: true}
);

export const createProjectEffect = createEffect(
    (
        actions$ = inject(Actions),
        service = inject(ProjectService),
    ) => {
        return actions$.pipe(
            ofType(CreateProjectAction),
            exhaustMap((action) => service.Create(action.payload).pipe(
                map(result => SetProjectAction({payload: result})),
                catchError(() => EMPTY)
            ))
        )
    },
    {functional: true}
);
