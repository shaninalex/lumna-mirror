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
import { catchError, exhaustMap, map, of, switchMap } from 'rxjs';
import { actionColumnSetList } from '@entities/column';
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
                this.taskApi.Create(action.board_id, action.data).pipe(
                    switchMap((task) => of(actionTaskUpsert({ task }))),
                    catchError((error) => of(actionTaskFailed({ error }))),
                ),
            ),
        ),
    );
}
