import {Actions, createEffect, ofType} from '@ngrx/effects';
import {inject} from '@angular/core';
import {
    TaskChangeStatusAction,
    TaskCreateAction,
    TaskListGetActions,
    TaskSetAction,
    TaskListSetActions, TaskUpdateAction, TaskPatchAction, TaskDeleteAction, TaskDeleteSuccessAction
} from './task.actions';
import {TaskService} from '../api/task.service';
import {exhaustMap, of, switchMap} from 'rxjs';
import {BoardViewApiService} from '@client/features/project/board-view-feature/api';


export const TasksGetEffect = createEffect(
    (
        actions$ = inject(Actions),
        api = inject(TaskService),
    ) => actions$.pipe(
        ofType(TaskListGetActions),
        exhaustMap((action) => {
                return api.List(action.projectId).pipe(
                    switchMap(data => of(TaskListSetActions({payload: data}))),
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
        ofType(TaskCreateAction),
        exhaustMap((action) => {
                return api.Create(action.projectId, action.payload).pipe(
                    switchMap(data => of(TaskSetAction({payload: data}))),
                )
            }
        )
    ),
    {functional: true, dispatch: true}
);

export const TasksChangeStatusEffect = createEffect(
    (
        actions$ = inject(Actions),
        api = inject(BoardViewApiService),
    ) => actions$.pipe(
        ofType(TaskChangeStatusAction),
        exhaustMap((action) => {
                return api.ChangeStatus(action.taskId, action.payload).pipe(
                    switchMap(data => of(TaskUpdateAction({payload: data}))),
                )
            }
        )
    ),
    {functional: true, dispatch: true}
);

export const TaskPatchEffect = createEffect(
    (
        actions$ = inject(Actions),
        api = inject(TaskService),
    ) => actions$.pipe(
        ofType(TaskPatchAction),
        exhaustMap((action) => {
                return api.Patch(action.taskId, action.payload).pipe(
                    switchMap(data => of(TaskUpdateAction({payload: data}))),
                )
            }
        )
    ),
    {functional: true, dispatch: true}
);

export const TaskDeleteEffect = createEffect(
    (
        actions$ = inject(Actions),
        api = inject(TaskService),
    ) => actions$.pipe(
        ofType(TaskDeleteAction),
        exhaustMap((action) => {
                return api.Delete(action.taskId).pipe(
                    switchMap(() => of(TaskDeleteSuccessAction({taskId: action.taskId}))),
                )
            }
        )
    ),
    {functional: true, dispatch: true}
);
