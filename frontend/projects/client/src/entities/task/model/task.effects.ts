import {Actions, createEffect, ofType} from '@ngrx/effects';
import {inject} from '@angular/core';
import {
    ChangeTaskStatusAction, ChangeTaskStatusSuccessAction,
    CreateTaskAction,
    GetTasksActions,
    SetTaskAction,
    SetTaskListActions
} from './task.actions';
import {TaskService} from '../api/task.service';
import {exhaustMap, of, switchMap} from 'rxjs';
import {BoardViewApiService} from '@client/features/project/board-view-feature/api';


export const TasksGetEffect = createEffect(
    (
        actions$ = inject(Actions),
        api = inject(TaskService),
    ) => actions$.pipe(
        ofType(GetTasksActions),
        exhaustMap((action) => {
                return api.List(action.projectId).pipe(
                    switchMap(data => of(SetTaskListActions({payload: data}))),
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
                return api.Create(action.projectId, action.payload).pipe(
                    switchMap(data => of(SetTaskAction({payload: data}))),
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
        ofType(ChangeTaskStatusAction),
        exhaustMap((action) => {
                return api.ChangeStatus(action.taskId, action.payload).pipe(
                    switchMap(data => of(ChangeTaskStatusSuccessAction({payload: data}))),
                )
            }
        )
    ),
    {functional: true, dispatch: true}
);
