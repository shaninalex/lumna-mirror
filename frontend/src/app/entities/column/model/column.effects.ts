import { Injectable, inject } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { ColumnsApi } from '../api/columns.api';
import {
    actionColumnCreate,
    actionColumnDeleteSuccess,
    actionColumnGetList,
    actionColumnPatch,
    actionColumnSetList,
    actionColumnUpsert,
} from './column.actions';
import { exhaustMap, map, of, switchMap, tap } from 'rxjs';
import { actionTaskSetTasks, TaskModel } from '@entities/task';

@Injectable()
export class ColumnEffects {
    private actions$ = inject(Actions);
    private columnsApi = inject(ColumnsApi);

    list_columns$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionColumnGetList),
            exhaustMap((action) =>
                this.columnsApi
                    .List(action.boardId)
                    .pipe(switchMap((columns) => of(actionColumnSetList({ columns })))),
            ),
        ),
    );

    create_column$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionColumnCreate),
            exhaustMap((action) =>
                this.columnsApi
                    .Create(action.boardId, action.data)
                    .pipe(switchMap((column) => of(actionColumnUpsert({ column })))),
            ),
        ),
    );

    patch_column$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionColumnPatch),
            exhaustMap((action) =>
                this.columnsApi
                    .Patch(action.columnId, action.data)
                    .pipe(switchMap((column) => of(actionColumnUpsert({ column })))),
            ),
        ),
    );

    delete_column$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionColumnPatch),
            exhaustMap((action) =>
                this.columnsApi
                    .Delete(action.columnId)
                    .pipe(
                        switchMap(() =>
                            of(actionColumnDeleteSuccess({ columnId: action.columnId })),
                        ),
                    ),
            ),
        ),
    );
}
