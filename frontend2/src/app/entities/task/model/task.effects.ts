import { inject, Injectable } from "@angular/core";
import { Actions, createEffect, ofType } from "@ngrx/effects";
import { catchError, of } from "rxjs";
import {
    actionTaskCreate,
    actionTaskCreateFailed,
    actionTaskCreateSuccess,
    actionTaskGetList,
    actionTaskSetList,
    actionTaskSetListFailed
} from "./task.actions";
import { TaskApi } from "@entities/task/api";
import { switchMap } from "rxjs/operators";
import { HttpErrorResponse } from "@angular/common/http";
import { fromErrorResponse } from "@shared/models";

@Injectable()
export class TaskEffects {
    private actions$ = inject(Actions);
    private api = inject(TaskApi);

    task_list$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionTaskGetList),
            switchMap((action) =>
                this.api.list(action.query).pipe(
                    switchMap((tasks) => of(actionTaskSetList({ tasks }))),
                    catchError((err: HttpErrorResponse) =>
                        of(
                            actionTaskSetListFailed({
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
            ofType(actionTaskCreate),
            switchMap((action) =>
                this.api.create(action.data).pipe(
                    switchMap((task) => of(actionTaskCreateSuccess({ task }))),
                    catchError((err: HttpErrorResponse) =>
                        of(
                            actionTaskCreateFailed({
                                errors: fromErrorResponse(err)
                            })
                        )
                    )
                )
            )
        )
    );
}
