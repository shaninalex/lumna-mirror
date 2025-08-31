import {Actions, createEffect, ofType} from "@ngrx/effects";
import {inject} from '@angular/core';
import {HttpClient, HttpErrorResponse} from '@angular/common/http';
import {Router} from '@angular/router';
import {GetSessionAction, SetSessionAction} from '@client/entities/session';
import {catchError, exhaustMap, of, switchMap} from 'rxjs';
import {Session} from '@ory/kratos-client';
import {environment} from '@client/environments/environment.development';

const urls = {
    session: `${environment.KRATOS_ROOT}/sessions/whoami`,
}

export const sessionGet = createEffect(
    (
        actions$ = inject(Actions),
        http = inject(HttpClient),
        router = inject(Router),
    ) => {
        return actions$.pipe(
            ofType(GetSessionAction.type),
            exhaustMap(() => http.get<Session>(urls.session, {withCredentials: true}).pipe(
                switchMap(data => of(SetSessionAction({session: data}))),
                catchError((error: HttpErrorResponse) => {
                    if (error.status === 404) {
                        router.navigate(['/404']);
                        return of({type: "[session] error"});
                    }
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
