import { inject, Injectable } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { TaskApi } from '../api/task.api';
import {
    actionTaskChange,
    actionTaskCreate,
    actionTaskFailed,
    actionTaskGetTaskById,
    actionTaskGetTasks,
    actionTaskSetTasks,
    actionTaskUpsert,
} from './task.actions';
import { catchError, exhaustMap, map, of, switchMap } from 'rxjs';
import { actionKanbanColumnsLoaded } from '@features/kanban-board';

@Injectable()
export class TaskEffects {
    private actions$ = inject(Actions);
    private taskApi = inject(TaskApi);

    // listen for kanban effect to load columns
    set_list$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionKanbanColumnsLoaded),
            map((action) => actionTaskSetTasks({ tasks: action.tasks })),
        ),
    );

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
                this.taskApi.Create(action.data).pipe(
                    switchMap((task) => of(actionTaskUpsert({ task }))),
                    catchError((error) => of(actionTaskFailed({ error }))),
                ),
            ),
        ),
    );

    task_get$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionTaskGetTaskById),
            exhaustMap((action) =>
                this.taskApi
                    .Get(action.task_id)
                    .pipe(switchMap((task) => of(actionTaskUpsert({ task })))),
            ),
        ),
    );

    task_change$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionTaskChange),
            exhaustMap((action) =>
                this.taskApi
                    .Patch(action.task_id, action.data)
                    .pipe(switchMap((task) => of(actionTaskUpsert({ task })))),
            ),
        )
    )
}
