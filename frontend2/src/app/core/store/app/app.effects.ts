import { Actions, createEffect, ofType } from "@ngrx/effects";
import { inject, Injectable } from "@angular/core";
import { exhaustMap, of } from "rxjs";

import { actionApplicationInit } from "./app.actions";
import { actionUserGet } from "@entities/user";

@Injectable()
export class ApplicationEffects {
    private actions$ = inject(Actions);

    appInit$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionApplicationInit),
            exhaustMap(() => of(actionUserGet()))
        )
    );
}
