import { createAction, props } from "@ngrx/store"
import { Comment } from "./comment.model"

export const CommentCreateAction = createAction("[comment] create", props<{ comment: Comment }>())

export const CommentsListAction = createAction("[comment] list", props<{ entity_id: number; entity_type: string }>())

export const CommentsSetAction = createAction("[comment] set", props<{ payload: Comment[] }>())
