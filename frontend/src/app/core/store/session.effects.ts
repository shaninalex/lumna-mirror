import { Actions, createEffect, ofType } from '@ngrx/effects';
import { inject, Injectable } from '@angular/core';
import { catchError, EMPTY, exhaustMap, map } from 'rxjs';
import { SessionApi } from './session.api';
import {
    actionSessionAuthenticatedSuccessfull,
    actionSessionAuthenticateStart,
    actionSessionLoggedOut,
    actionSessionLoggingOut,
} from './session.actions';
import { actionUserClear, actionUserSet } from '@entities/user';
import { Router } from '@angular/router';

@Injectable()
export class SessionEffects {
    private actions$ = inject(Actions);
    private sessionApi = inject(SessionApi);
    private router = inject(Router);

    authenticate_start$ = createEffect(() => {
        return this.actions$.pipe(
            ofType(actionSessionAuthenticateStart.type),
            exhaustMap((action) =>
                this.sessionApi.login(action.email, action.password).pipe(
                    map((user) => actionSessionAuthenticatedSuccessfull({ user: user })),
                    catchError(() => EMPTY), // TODO: handle auth errors
                ),
            ),
        );
    });

    logging_out_init$ = createEffect(() => {
        return this.actions$.pipe(
            ofType(actionSessionLoggingOut.type),
            exhaustMap(() =>
                this.sessionApi.logout().pipe(
                    map(() => actionSessionLoggedOut()),
                    catchError(() => EMPTY),
                ),
            ),
        );
    });

    authenticated_successfull$ = createEffect(() => {
        return this.actions$.pipe(
            ofType(actionSessionAuthenticatedSuccessfull.type),
            map((action) => actionUserSet({ user: action.user })),
        );
    });

    authenticated_successfull_redirect$ = createEffect(
        () => {
            return this.actions$.pipe(
                ofType(actionSessionAuthenticatedSuccessfull.type),
                map(() => this.router.navigateByUrl('/')),
            );
        },
        { dispatch: false },
    );

    logging_out_completed$ = createEffect(() => {
        return this.actions$.pipe(
            ofType(actionSessionLoggedOut.type),
            map((action) => actionUserClear()),
        );
    });

    logging_out_completed_redirect$ = createEffect(
        () => {
            return this.actions$.pipe(
                ofType(actionSessionLoggedOut.type),
                map(() => this.router.navigateByUrl('/auth/login')),
            );
        },
        { dispatch: false },
    );
}
