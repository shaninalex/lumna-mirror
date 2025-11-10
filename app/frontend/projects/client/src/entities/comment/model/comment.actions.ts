import { createAction, props } from "@ngrx/store"
import { Comment } from "./comment.model"

export const CommentCreateAction = createAction("[comment] create", props<{ comment: Comment }>())

export const CommentCreateSuccessAction = createAction("[comment] create success", props<{ payload: Comment }>())

export const CommentsListAction = createAction("[comment] list", props<{ entity_id: number; entity_type: string }>())
