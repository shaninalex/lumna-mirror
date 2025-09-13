import {Actions, createEffect, ofType} from '@ngrx/effects';
import {inject} from '@angular/core';
import {catchError, EMPTY, exhaustMap, map} from 'rxjs';
import {UserService} from '@client/entities/user';
import {GetUserAction, SetUserAction, UpdateUserSettings} from '@client/entities/user/model/user.actions';

export const GetUserEffect = createEffect(
    (
        actions$ = inject(Actions),
        service = inject(UserService),
    ) => {
        return actions$.pipe(
            ofType(GetUserAction),
            exhaustMap(() => service.getUser().pipe(
                map(result => SetUserAction({payload: result})),
                catchError(() => EMPTY)
            ))
        )
    },
    {functional: true, dispatch: true}
);

export const UpdateUserSettingsEffect = createEffect(
    (
        actions$ = inject(Actions),
        service = inject(UserService),
    ) => {
        return actions$.pipe(
            ofType(UpdateUserSettings),
            exhaustMap((action) => service.updateUserSettings(action.payload).pipe(
                map(result => SetUserAction({payload: result})),
                catchError(() => EMPTY)
            ))
        )
    },
    {functional: true, dispatch: true}
);
