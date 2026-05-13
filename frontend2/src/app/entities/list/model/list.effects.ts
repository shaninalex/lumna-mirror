import { inject, Injectable } from "@angular/core";
import { Actions, createEffect, ofType } from "@ngrx/effects";
import { exhaustMap, of } from "rxjs";
import { actionListGetList, actionListSetList } from "./list.actions";

@Injectable()
export class ListEffects {
    private actions$ = inject(Actions);

    list_list$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionListGetList),
            exhaustMap(() => {
                // TODO: api
                return of(actionListSetList({ list: [] }));
            })
        )
    );
}
