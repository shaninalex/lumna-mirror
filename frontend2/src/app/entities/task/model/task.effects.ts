import { inject, Injectable } from "@angular/core";
import { Actions, createEffect, ofType } from "@ngrx/effects";
import { exhaustMap, of } from "rxjs";
import { actionTaskGetList, actionTaskSetList } from "./task.actions";

@Injectable()
export class TaskEffects {
    private actions$ = inject(Actions);

    task_list$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionTaskGetList),
            exhaustMap(() => {
                // TODO: api
                return of(actionTaskSetList({ list: [] }));
            })
        )
    );
}
