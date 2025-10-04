import {Actions, createEffect, ofType} from '@ngrx/effects';
import {inject} from '@angular/core';
import {CreateTaskAction, GetTasksActions, SetTaskAction, SetTasksActions} from './task.actions';
import {TaskService} from '../api/task.service';
import {exhaustMap, of, switchMap} from 'rxjs';


export const TasksGetEffect = createEffect(
    (
        actions$ = inject(Actions),
        api = inject(TaskService),
    ) => actions$.pipe(
        ofType(GetTasksActions),
        exhaustMap((action) => {
                return api.List(action.projectId).pipe(
                    switchMap(data => of(SetTasksActions({payload: data}))),
                )
            }
        )
    ),
    {functional: true, dispatch: true}
);

export const TasksCreateEffect = createEffect(
    (
        actions$ = inject(Actions),
        api = inject(TaskService),
    ) => actions$.pipe(
        ofType(CreateTaskAction),
        exhaustMap((action) => {
                return api.Create(action.payload).pipe(
                    switchMap(data => of(SetTaskAction({payload: data}))),
                )
            }
        )
    ),
    {functional: true, dispatch: true}
);
