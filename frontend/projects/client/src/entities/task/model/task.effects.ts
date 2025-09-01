import {Actions, createEffect, ofType} from '@ngrx/effects';
import {inject} from '@angular/core';
import {GetTasksActions, SetTasksActions} from './task.actions';
import {TaskService} from '../api/task.service';
import {exhaustMap, of, switchMap} from 'rxjs';


export const GetTasksEffect = createEffect(
    (
        actions$ = inject(Actions),
        api = inject(TaskService),
    ) => actions$.pipe(
        ofType(GetTasksActions.type),
        exhaustMap((action) => {
                return api.List(action.projectCode).pipe(
                    switchMap(data => of(SetTasksActions({payload: data}))),
                )
            }
        )
    ),
    {functional: true, dispatch: true}
);
