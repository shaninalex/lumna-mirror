import { inject, Injectable } from "@angular/core";
import { Actions, createEffect, ofType } from "@ngrx/effects";
import { exhaustMap, of } from "rxjs";
import { actionProjectGetList, actionProjectSetList } from "./project.actions";

@Injectable()
export class ProjectEffects {
    private actions$ = inject(Actions);

    project_list$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionProjectGetList),
            exhaustMap(() => {
                // TODO: api
                return of(actionProjectSetList({ list: [] }));
            })
        )
    );
}
