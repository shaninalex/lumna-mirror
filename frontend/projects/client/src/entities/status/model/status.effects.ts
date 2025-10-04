import {Actions, createEffect, ofType} from '@ngrx/effects';
import {inject} from '@angular/core';
import {exhaustMap, of, switchMap} from 'rxjs';
import {
    CreateStatusAction, DeleteStatusAction, DeleteStatusSuccessAction,
    GetStatusListActions, PatchStatusAction, PatchStatusSuccessAction,
    SetStatusAction,
    SetStatusListActions,
    StatusService
} from '@client/entities/status';


export const statusGetEffect = createEffect(
    (
        actions$ = inject(Actions),
        api = inject(StatusService),
    ) => actions$.pipe(
        ofType(GetStatusListActions),
        exhaustMap((action) => {
                return api.List(action.projectId).pipe(
                    switchMap(data => of(SetStatusListActions({payload: data}))),
                )
            }
        )
    ),
    {functional: true, dispatch: true}
);

export const statusCreateEffect = createEffect(
    (
        actions$ = inject(Actions),
        api = inject(StatusService),
    ) => actions$.pipe(
        ofType(CreateStatusAction),
        exhaustMap((action) => {
                return api.Create(action.projectId, action.payload).pipe(
                    switchMap(data => of(SetStatusAction({payload: data}))),
                )
            }
        )
    ),
    {functional: true, dispatch: true}
);

export const statusPatchEffect = createEffect(
    (
        actions$ = inject(Actions),
        api = inject(StatusService),
    ) => actions$.pipe(
        ofType(PatchStatusAction),
        exhaustMap((action) => {
                return api.Patch(action.projectId, action.statusId, action.payload).pipe(
                    switchMap(data => of(PatchStatusSuccessAction({payload: data}))),
                )
            }
        )
    ),
    {functional: true, dispatch: true}
);

export const statusDeleteEffect = createEffect(
    (
        actions$ = inject(Actions),
        api = inject(StatusService),
    ) => actions$.pipe(
        ofType(DeleteStatusAction),
        exhaustMap((action) => {
                return api.Delete(action.projectId, action.statusId).pipe(
                    switchMap(() => of(DeleteStatusSuccessAction({projectId: action.projectId, statusId: action.statusId}))),
                )
            }
        )
    ),
    {functional: true, dispatch: true}
)
