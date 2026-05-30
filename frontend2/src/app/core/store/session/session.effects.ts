import { Actions, createEffect, ofType } from "@ngrx/effects";
import { inject, Injectable } from "@angular/core";
import { catchError, EMPTY, exhaustMap, map, of } from "rxjs";
import { SessionApi } from "./session.api";
import {
    actionLoginFailed,
    actionSessionAuthenticated,
    actionSessionAuthenticatedSuccessfull,
    actionSessionAuthenticateStart,
    actionSessionFailed,
    actionSessionLoggedOut,
    actionSessionLoggingOut
} from "./session.actions";
import { actionUserClear, actionUserSet } from "@entities/user";
import { Router } from "@angular/router";
import { HttpErrorResponse } from "@angular/common/http";
import { fromErrorResponse } from "@shared/models";

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
                    map((user) =>
                        actionSessionAuthenticatedSuccessfull({ user: user })
                    ),
                    catchError((err: HttpErrorResponse) =>
                        of(
                            actionLoginFailed({
                                errors: fromErrorResponse(err)
                            })
                        )
                    )
                )
            )
        );
    });

    logging_out_init$ = createEffect(() => {
        return this.actions$.pipe(
            ofType(actionSessionLoggingOut.type),
            exhaustMap(() =>
                this.sessionApi.logout().pipe(
                    map(() => actionSessionLoggedOut()),
                    catchError(() => EMPTY)
                )
            )
        );
    });

    authenticated_successfull$ = createEffect(() => {
        return this.actions$.pipe(
            ofType(actionSessionAuthenticatedSuccessfull.type),
            map((action) => actionUserSet({ user: action.user }))
        );
    });

    user_set_marks_authenticated$ = createEffect(() => {
        return this.actions$.pipe(
            ofType(actionUserSet.type),
            map(() => actionSessionAuthenticated())
        );
    });

    authenticated_successfull_redirect$ = createEffect(
        () => {
            return this.actions$.pipe(
                ofType(actionSessionAuthenticatedSuccessfull),
                map(() => this.router.navigateByUrl("/app"))
            );
        },
        { dispatch: false }
    );

    logging_out_completed$ = createEffect(() => {
        return this.actions$.pipe(
            ofType(actionSessionLoggedOut.type),
            map(() => actionUserClear())
        );
    });

    logging_out_completed_redirect$ = createEffect(
        () => {
            return this.actions$.pipe(
                ofType(actionSessionLoggedOut.type),
                map(() => this.router.navigateByUrl("/auth/login"))
            );
        },
        { dispatch: false }
    );

    authenticate_failed$ = createEffect(
        () => {
            return this.actions$.pipe(
                ofType(actionSessionFailed)
                // map(() => this.router.navigateByUrl("/auth/login"))
            );
        },
        { dispatch: false }
    );
}
