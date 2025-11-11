import { Actions, createEffect, ofType } from "@ngrx/effects"
import { inject } from "@angular/core"
import { CommentCreateAction, CommentsListAction, CommentsSetAction } from "./comment.actions"
import { CommentService } from "../api/comment.service"
import { exhaustMap, of, switchMap } from "rxjs"

export const CommentCreateEffect = createEffect(
	(actions$ = inject(Actions), api = inject(CommentService)) =>
		actions$.pipe(
			ofType(CommentCreateAction),
			exhaustMap(action => api.Create(action.comment).pipe(switchMap(data => of(CommentsSetAction({ payload: [data] })))))
		),
	{ functional: true, dispatch: true }
)

export const CommentListEffect = createEffect(
	(actions$ = inject(Actions), api = inject(CommentService)) =>
		actions$.pipe(
			ofType(CommentsListAction),
			exhaustMap(action =>
				api.List(action.entity_id, action.entity_type).pipe(switchMap(data => of(CommentsSetAction({ payload: data }))))
			)
		),
	{ functional: true, dispatch: true }
)
