import { Injectable, inject } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { TaskApi } from '../api/task.api';
import {
    actionTaskCreate,
    actionTaskFailed,
    actionTaskGetTasks,
    actionTaskSetTasks,
    actionTaskUpsert,
} from './task.actions';
import { catchError, exhaustMap, map, of, switchMap, tap } from 'rxjs';
import { TaskModel } from './task.model';
import { actionColumnSetList } from '@entities/column';

@Injectable()
export class TaskEffects {
    private actions$ = inject(Actions);
    private taskApi = inject(TaskApi);

    task_list$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionTaskGetTasks),
            exhaustMap((action) =>
                this.taskApi
                    .List(action.board_id)
                    .pipe(switchMap((tasks) => of(actionTaskSetTasks({ tasks })))),
            ),
        ),
    );

    task_create$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionTaskCreate),
            exhaustMap((action) =>
                this.taskApi.Create(action.board_id, action.data).pipe(
                    switchMap((task) => of(actionTaskUpsert({ task }))),
                    catchError((error) => of(actionTaskFailed({ error }))),
                ),
            ),
        ),
    );

    add_column_tasks$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionColumnSetList),
            map((action) => {
                const tasks: TaskModel[] = [];
                for (let ci = 0; ci < action.columns.length; ci++) {
                    for (let ti = 0; ti < action.columns[ci].tasks.length; ti++) {
                        tasks.push(action.columns[ci].tasks[ti]);
                    }
                }
                return actionTaskSetTasks({ tasks });
            }),
        ),
    );
}
