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
import { catchError, exhaustMap, of, switchMap } from 'rxjs';

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
}
