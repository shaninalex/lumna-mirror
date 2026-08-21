import { inject, Injectable } from "@angular/core";
import { Actions, createEffect, ofType } from "@ngrx/effects";
import { catchError, of } from "rxjs";
import { switchMap } from "rxjs/operators";
import type { HttpErrorResponse } from "@angular/common/http";
import { fromErrorResponse } from "@shared/models";
import { actionsStatuses } from "./status.actions";
import { StatusApi } from "../api";

@Injectable()
export class StatusEffects {
    private actions$ = inject(Actions);
    private api = inject(StatusApi);

    statuses$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionsStatuses.loadByListId),
            switchMap((action) =>
                this.api.list(action.list_id).pipe(
                    switchMap((statuses) => of(actionsStatuses.loadByListIdSuccess({ statuses }))),
                    catchError((err: HttpErrorResponse) =>
                        of(
                            actionsStatuses.loadByListIdFailed({
                                errors: fromErrorResponse(err)
                            })
                        )
                    )
                )
            )
        )
    );

    task_create$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionsStatuses.create),
            switchMap((action) =>
                this.api.create(action.payload).pipe(
                    switchMap((status) => of(actionsStatuses.createSuccess({ status }))),
                    catchError((err: HttpErrorResponse) =>
                        of(
                            actionsStatuses.createFailed({
                                errors: fromErrorResponse(err)
                            })
                        )
                    )
                )
            )
        )
    );
}
