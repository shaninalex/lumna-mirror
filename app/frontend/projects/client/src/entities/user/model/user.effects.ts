import {Actions, createEffect, ofType} from '@ngrx/effects';
import {inject} from '@angular/core';
import {catchError, EMPTY, exhaustMap, map, of} from 'rxjs';
import {UserService} from '@client/entities/user';
import {
    UserGetAction,
    UserSetAction,
    UserUpdateSettingsAction
} from '@client/entities/user/model/user.actions';

export const UserGetEffect = createEffect(
    (
        actions$ = inject(Actions),
        service = inject(UserService),
    ) => {
        return actions$.pipe(
            ofType(UserGetAction),
            exhaustMap(() =>
                service.getUser().pipe(
                    map(user => UserSetAction({ payload: user })),
                    catchError((err) => {
                        if (err.status === 401) {
                            return of(UserSetAction({ payload: null }));
                        }
                        return of({ type: "api_error", error: err });
                    })
                )
            )
        );
    },
    { functional: true, dispatch: true }
);

export const UserUpdateSettingsEffect = createEffect(
    (
        actions$ = inject(Actions),
        service = inject(UserService),
    ) => {
        return actions$.pipe(
            ofType(UserUpdateSettingsAction),
            exhaustMap((action) => service.updateUserSettings(action.payload).pipe(
                map(result => UserSetAction({payload: result})),
                catchError(() => EMPTY)
            ))
        )
    },
    {functional: true, dispatch: true}
);
