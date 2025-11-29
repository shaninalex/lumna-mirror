import { createEntityAdapter, EntityAdapter, EntityState } from "@ngrx/entity"
import { createReducer, on } from "@ngrx/store"
import { CommentSetAction } from "./comment.actions"
import { Comment } from "./comment.model"

export interface CommentsState extends EntityState<Comment> {}
export const commentsAdapter: EntityAdapter<Comment> = createEntityAdapter<Comment>()
export const commentsReducer = createReducer(
	commentsAdapter.getInitialState(),
	on(CommentSetAction, (state, action) => commentsAdapter.addMany(action.payload, state))
)
