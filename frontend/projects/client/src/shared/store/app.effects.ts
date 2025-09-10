import {Actions, createEffect, ofType} from '@ngrx/effects';
import {inject} from '@angular/core';
import {EMPTY, exhaustMap} from 'rxjs';
import {AppErrorAction} from '@client/shared/store/app.actions';
import {Router} from '@angular/router';

export const HandleError = createEffect(
    (
        actions$ = inject(Actions),
        router= inject(Router),
    ) => actions$.pipe(
        ofType(AppErrorAction.type),
        exhaustMap((action) => {
                console.log("HandleError: ", action)
                switch (action.err.key) {
                    case "org_not_attached":
                        router.navigate(['/set-organization']);
                        break
                }
                return EMPTY
            }
        )
    ),
    {functional: true, dispatch: false}
);
