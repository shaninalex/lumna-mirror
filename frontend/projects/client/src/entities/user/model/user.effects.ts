import {Actions, createEffect, ofType} from '@ngrx/effects';
import {inject} from '@angular/core';
import {catchError, EMPTY, exhaustMap, map, of} from 'rxjs';
import {UserService} from '@client/entities/user';
import {
    GetUserAction,
    SetUserAction,
    UpdateUserSettingsAction
} from '@client/entities/user/model/user.actions';

export const GetUserEffect = createEffect(
    (
        actions$ = inject(Actions),
        service = inject(UserService),
    ) => {
        return actions$.pipe(
            ofType(GetUserAction),
            exhaustMap(() =>
                service.getUser().pipe(
                    map(user => SetUserAction({ payload: user })),
                    catchError((err) => {
                        if (err.status === 401) {
                            return of(SetUserAction({ payload: null }));
                        }
                        return of({ type: "api_error", error: err });
                    })
                )
            )
        );
    },
    { functional: true, dispatch: true }
);

export const UpdateUserSettingsEffect = createEffect(
    (
        actions$ = inject(Actions),
        service = inject(UserService),
    ) => {
        return actions$.pipe(
            ofType(UpdateUserSettingsAction),
            exhaustMap((action) => service.updateUserSettings(action.payload).pipe(
                map(result => SetUserAction({payload: result})),
                catchError(() => EMPTY)
            ))
        )
    },
    {functional: true, dispatch: true}
);
