import { Actions, createEffect, ofType } from "@ngrx/effects";
import { inject, Injectable } from "@angular/core";
import { tap } from "rxjs";

import { actionApplicationInit } from "./app.actions";

@Injectable()
export class ApplicationEffects {
    private actions$ = inject(Actions);

    appInit$ = createEffect(
        () =>
            this.actions$.pipe(
                ofType(actionApplicationInit),
                tap(() => {
                    console.log("application init");
                })
                // TODO:
                // switchMap(() =>
                //     this.authService.getUser().pipe(
                //         map((user) => AppActions.authSuccess({ user })),
                //         catchError(() => of(AppActions.authFailed()))
                //     )
                // )
            ),
        { dispatch: false }
    );
}
