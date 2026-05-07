import { inject, Injectable } from "@angular/core";
import { ColumnsApi } from "@entities/column/api/columns.api";
import { Actions, createEffect, ofType } from "@ngrx/effects";
import {
    actionKanbanLoadColumns,
    actionKanbanColumnsLoaded
} from "./shared.actions";
import { exhaustMap, of, switchMap } from "rxjs";
import { ColumnModel } from "@entities/column";
import { TaskApi } from "@entities/task/api/task.api";
import { actionTaskSetTasks } from "@entities/task";

@Injectable()
export class KanbanEffects {
    private actions$ = inject(Actions);
    private columnsApi = inject(ColumnsApi);
    private tasksApi = inject(TaskApi);

    load_columns$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionKanbanLoadColumns),
            exhaustMap((action) =>
                this.columnsApi.List(action.boardId).pipe(
                    switchMap((apiColumns) => {
                        const columns: ColumnModel[] = [];
                        for (const col of apiColumns) {
                            columns.push({
                                ...col,
                                created_at: new Date(col.created_at),
                                updated_at: new Date(col.updated_at)
                            });
                        }
                        return of(
                            actionKanbanColumnsLoaded({
                                boardId: action.boardId,
                                columns
                            })
                        );
                    })
                )
            )
        )
    );

    load_tasks$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionKanbanColumnsLoaded),
            exhaustMap((action) =>
                this.tasksApi
                    .List(action.boardId)
                    .pipe(
                        switchMap((tasks) => of(actionTaskSetTasks({ tasks })))
                    )
            )
        )
    );
}
