import { inject, Injectable } from "@angular/core";
import { Actions, createEffect, ofType } from "@ngrx/effects";
import { catchError, of } from "rxjs";
import { switchMap } from "rxjs/operators";
import type { HttpErrorResponse } from "@angular/common/http";
import { fromErrorResponse } from "@shared/models";
import { statusActions } from "./status.actions";
import { StatusApi } from "../api";

@Injectable()
export class StatusEffects {
    private actions$ = inject(Actions);
    private api = inject(StatusApi);

    statuses$ = createEffect(() =>
        this.actions$.pipe(
            ofType(statusActions.loadByListId),
            switchMap((action) =>
                this.api.list(action.list_id).pipe(
                    switchMap((statuses) => of(statusActions.loadByListIdSuccess({ statuses }))),
                    catchError((err: HttpErrorResponse) =>
                        of(
                            statusActions.loadByListIdFailed({
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
            ofType(statusActions.create),
            switchMap((action) =>
                this.api.create(action.payload).pipe(
                    switchMap((status) => of(statusActions.createSuccess({ status }))),
                    catchError((err: HttpErrorResponse) =>
                        of(
                            statusActions.createFailed({
                                errors: fromErrorResponse(err)
                            })
                        )
                    )
                )
            )
        )
    );
}
