import { createAction, props } from "@ngrx/store"
import { Comment } from "./comment.model"

export const CommentCreateAction = createAction("[comment] create", props<{ comment: Comment }>())

export const CommentListAction = createAction("[comment] list", props<{ entity_id: number; entity_type: string }>())

export const CommentSetAction = createAction("[comment] set", props<{ payload: Comment[] }>())
