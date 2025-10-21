import {Actions, createEffect, ofType} from '@ngrx/effects';
import {inject} from '@angular/core';
import {EMPTY, exhaustMap} from 'rxjs';
import {AppErrorAction} from '@client/shared/store/app.actions';

export const HandleError = createEffect(
    (
        actions$ = inject(Actions),
    ) => actions$.pipe(
        ofType(AppErrorAction.type),
        exhaustMap((action) => {
                return EMPTY
            }
        )
    ),
    {functional: true, dispatch: false}
);
