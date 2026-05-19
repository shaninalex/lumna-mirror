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
                ofType(actionApplicationInit)
                // tap((action) => console.log(action))
                // exhaustMap(() => of(actionUserGet()))
            ),
        { dispatch: false }
    );
}
