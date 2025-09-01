import {Actions, createEffect, ofType} from '@ngrx/effects';
import {inject} from '@angular/core';
import {HttpClient, HttpErrorResponse} from '@angular/common/http';
import {Router} from '@angular/router';
import {GetProjectsAction, SetProjectsAction} from '@client/entities/project/model/project.actions';
import {catchError, exhaustMap, of, switchMap} from 'rxjs';
import {ProjectService} from '@client/entities/project/api/project.service';

export const getProjectsEffect = createEffect(
    (
        actions$ = inject(Actions),
        http = inject(ProjectService),
        router = inject(Router),
    ) => {
        return actions$.pipe(
            ofType(GetProjectsAction.type),
            exhaustMap(() => http.List().pipe(
                switchMap(data => of(SetProjectsAction({payload: data}))),
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
