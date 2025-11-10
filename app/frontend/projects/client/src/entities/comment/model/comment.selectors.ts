import { createFeatureSelector, createSelector } from "@ngrx/store"
import { CommentsState, commentsAdapter } from "./comment.reducer"

export const selectCommentsFeature = createFeatureSelector<CommentsState>("comments")
export const commentsSelectors = commentsAdapter.getSelectors(selectCommentsFeature)
export const selectAllComments = commentsSelectors.selectAll

export const selectComments = (entity_id: number, entity_type: string) =>
	createSelector(selectAllComments, tasks => tasks.filter(t => t.entity_id === entity_id && entity_type === entity_type))
