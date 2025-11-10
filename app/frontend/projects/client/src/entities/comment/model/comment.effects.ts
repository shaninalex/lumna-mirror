import { Actions, createEffect, ofType } from "@ngrx/effects"
import { inject } from "@angular/core"
import { CommentCreateAction, CommentCreateSuccessAction } from "./comment.actions"
import { CommentService } from "../api/comment.service"
import { exhaustMap, of, switchMap } from "rxjs"

export const CommentCreateEffect = createEffect(
	(actions$ = inject(Actions), api = inject(CommentService)) =>
		actions$.pipe(
			ofType(CommentCreateAction),
			exhaustMap(action => api.Create(action.comment).pipe(switchMap(data => of(CommentCreateSuccessAction({ payload: data })))))
		),
	{ functional: true, dispatch: true }
)
