import { inject, Injectable } from "@angular/core";
import { Actions, createEffect, ofType } from "@ngrx/effects";
import { catchError, of, withLatestFrom } from "rxjs";
import {
    actionSprintCreate,
    actionSprintCreateFailed,
    actionSprintCreateSuccess,
    actionSprintGetList,
    actionSprintSetList,
    actionSprintSetListFailed
} from "./sprint.actions";
import { SprintApi } from "@entities/sprint/api";
import { filter, switchMap } from "rxjs/operators";
import { HttpErrorResponse } from "@angular/common/http";
import { fromErrorResponse } from "@shared/models";
import { Store } from "@ngrx/store";
import { selectProjectCurrentId } from "@entities/project";

@Injectable()
export class SprintEffects {
    private actions$ = inject(Actions);
    private api = inject(SprintApi);
    private store = inject(Store);

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
            withLatestFrom(this.store.select(selectProjectCurrentId)),
            filter(([, projectId]) => projectId !== undefined),
            switchMap(([action, projectId]) =>
                this.api.create({ ...action.data, project_id: projectId! }).pipe(
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
