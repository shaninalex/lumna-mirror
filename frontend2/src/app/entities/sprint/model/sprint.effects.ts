import { inject, Injectable } from "@angular/core";
import { Actions, createEffect, ofType } from "@ngrx/effects";
import { catchError, of } from "rxjs";
import {
    actionSprintCreate,
    actionSprintCreateFailed,
    actionSprintCreateSuccess,
    actionSprintGetList,
    actionSprintSetList,
    actionSprintSetListFailed
} from "./sprint.actions";
import { SprintApi } from "@entities/sprint/api";
import { switchMap } from "rxjs/operators";
import { HttpErrorResponse } from "@angular/common/http";
import { fromErrorResponse } from "@shared/models";

@Injectable()
export class SprintEffects {
    private actions$ = inject(Actions);
    private api = inject(SprintApi);

    sprint_list$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionSprintGetList),
            switchMap((action) =>
                this.api.list(action.query).pipe(
                    switchMap((sprints) =>
                        of(actionSprintSetList({ sprints }))
                    ),
                    catchError((err: HttpErrorResponse) =>
                        of(
                            actionSprintSetListFailed({
                                errors: fromErrorResponse(err)
                            })
                        )
                    )
                )
            )
        )
    );

    sprint_Create$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionSprintCreate),
            switchMap((action) =>
                this.api.create(action.data).pipe(
                    switchMap((sprint) =>
                        of(actionSprintCreateSuccess({ sprint }))
                    ),
                    catchError((err: HttpErrorResponse) =>
                        of(
                            actionSprintCreateFailed({
                                errors: fromErrorResponse(err)
                            })
                        )
                    )
                )
            )
        )
    );
}
