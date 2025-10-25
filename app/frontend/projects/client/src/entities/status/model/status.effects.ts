import { Actions, createEffect, ofType } from "@ngrx/effects"
import { inject } from "@angular/core"
import { exhaustMap, of, switchMap } from "rxjs"
import {
	StatusCreateAction,
	StatusDeleteAction,
	StatusDeleteSuccessAction,
	StatusListGetActions,
	StatusListSetActions,
	StatusPatchAction,
	StatusPatchSortAction,
	StatusPatchSortSuccessAction,
	StatusPatchSuccessAction,
	StatusService,
	StatusSetAction,
} from "@client/entities/status"

export const StatusGetEffect = createEffect(
	(actions$ = inject(Actions), api = inject(StatusService)) =>
		actions$.pipe(
			ofType(StatusListGetActions),
			exhaustMap(action => {
				return api.List(action.projectId).pipe(switchMap(data => of(StatusListSetActions({ payload: data }))))
			})
		),
	{ functional: true, dispatch: true }
)

export const StatusCreateEffect = createEffect(
	(actions$ = inject(Actions), api = inject(StatusService)) =>
		actions$.pipe(
			ofType(StatusCreateAction),
			exhaustMap(action => {
				return api.Create(action.projectId, action.payload).pipe(switchMap(data => of(StatusSetAction({ payload: data }))))
			})
		),
	{ functional: true, dispatch: true }
)

export const StatusPatchEffect = createEffect(
	(actions$ = inject(Actions), api = inject(StatusService)) =>
		actions$.pipe(
			ofType(StatusPatchAction),
			exhaustMap(action => {
				return api
					.Patch(action.projectId, action.statusId, action.payload)
					.pipe(switchMap(data => of(StatusPatchSuccessAction({ payload: data }))))
			})
		),
	{ functional: true, dispatch: true }
)

export const StatusDeleteEffect = createEffect(
	(actions$ = inject(Actions), api = inject(StatusService)) =>
		actions$.pipe(
			ofType(StatusDeleteAction),
			exhaustMap(action => {
				return api
					.Delete(action.projectId, action.statusId)
					.pipe(switchMap(() => of(StatusDeleteSuccessAction({ projectId: action.projectId, statusId: action.statusId }))))
			})
		),
	{ functional: true, dispatch: true }
)

export const StatusSortPatchEffect = createEffect(
	(actions$ = inject(Actions), api = inject(StatusService)) =>
		actions$.pipe(
			ofType(StatusPatchSortAction),
			exhaustMap(action => {
				return api
					.StatusSort(action.projectId, action.payload)
					.pipe(switchMap(data => of(StatusPatchSortSuccessAction({ payload: data }))))
			})
		),
	{ functional: true, dispatch: true }
)
