import { inject, Injectable } from '@angular/core';
import { ColumnsApi } from '@entities/column/api/columns.api';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { actionKanbanLoadColumns, actionKanbanColumnsLoaded } from './shared.actions';
import { exhaustMap, of, switchMap } from 'rxjs';
import { ColumnModel } from '@entities/column';
import { TaskModel } from '@entities/task';

@Injectable()
export class KanbanEffects {
    private actions$ = inject(Actions);
    private columnsApi = inject(ColumnsApi);

    load_columns$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionKanbanLoadColumns),
            exhaustMap((action) =>
                this.columnsApi.List(action.boardId).pipe(
                    switchMap((apiColumns) => {
                        const columns: ColumnModel[] = [];
                        const tasks: TaskModel[] = [];
                        for (const col of apiColumns) {
                            columns.push({
                                ...col,
                                created_at: new Date(col.created_at),
                                updated_at: new Date(col.updated_at),
                            });

                            for (const task of col.tasks) {
                                tasks.push({
                                    ...task,
                                    created_at: new Date(task.created_at),
                                    updated_at: new Date(task.updated_at),
                                });
                            }
                        }
                        return of(actionKanbanColumnsLoaded({ columns, tasks }));
                    }),
                ),
            ),
        ),
    );
}
