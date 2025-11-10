import { Actions, createEffect, ofType } from "@ngrx/effects"
import { inject } from "@angular/core"
import {
	TaskChangeStatusAction,
	TaskCreateAction,
	TaskCreateCommentAction,
	TaskCreateCommentSuccessAction,
	TaskDeleteAction,
	TaskDeleteSuccessAction,
	TaskListGetActions,
	TaskListSetActions,
	TaskPatchAction,
	TaskSetAction,
	TaskUpdateAction,
} from "./task.actions"
import { TaskService } from "../api/task.service"
import { exhaustMap, of, switchMap } from "rxjs"
import { BoardViewApiService } from "@client/features/project/board-view-feature/api"
import { Comment } from "./task.model"

export const TasksGetEffect = createEffect(
	(actions$ = inject(Actions), api = inject(TaskService)) =>
		actions$.pipe(
			ofType(TaskListGetActions),
			exhaustMap(action => {
				return api.List(action.projectId).pipe(switchMap(data => of(TaskListSetActions({ payload: data }))))
			})
		),
	{ functional: true, dispatch: true }
)

export const TasksCreateEffect = createEffect(
	(actions$ = inject(Actions), api = inject(TaskService)) =>
		actions$.pipe(
			ofType(TaskCreateAction),
			exhaustMap(action => {
				return api.Create(action.projectId, action.payload).pipe(switchMap(data => of(TaskSetAction({ payload: data }))))
			})
		),
	{ functional: true, dispatch: true }
)

export const TasksChangeStatusEffect = createEffect(
	(actions$ = inject(Actions), api = inject(BoardViewApiService)) =>
		actions$.pipe(
			ofType(TaskChangeStatusAction),
			exhaustMap(action => {
				return api.ChangeStatus(action.taskId, action.payload).pipe(switchMap(data => of(TaskUpdateAction({ payload: data }))))
			})
		),
	{ functional: true, dispatch: true }
)

export const TaskPatchEffect = createEffect(
	(actions$ = inject(Actions), api = inject(TaskService)) =>
		actions$.pipe(
			ofType(TaskPatchAction),
			exhaustMap(action => {
				return api.Patch(action.taskId, action.payload).pipe(switchMap(data => of(TaskUpdateAction({ payload: data }))))
			})
		),
	{ functional: true, dispatch: true }
)

export const TaskDeleteEffect = createEffect(
	(actions$ = inject(Actions), api = inject(TaskService)) =>
		actions$.pipe(
			ofType(TaskDeleteAction),
			exhaustMap(action => {
				return api.Delete(action.taskId).pipe(switchMap(() => of(TaskDeleteSuccessAction({ taskId: action.taskId }))))
			})
		),
	{ functional: true, dispatch: true }
)

export const TaskCommentCreateEffect = createEffect(
	(actions$ = inject(Actions), api = inject(TaskService)) =>
		actions$.pipe(
			ofType(TaskCreateCommentAction),
			exhaustMap(action => {
				const comment: Comment = {
					id: 0,
					task_id: action.taskId,
					user_id: 0,
					content: action.message,
					created_at: new Date(),
				}
				return api
					.CreateComment(action.taskId, comment)
					.pipe(switchMap(data => of(TaskCreateCommentSuccessAction({ payload: data }))))
			})
		),
	{ functional: true, dispatch: true }
)
