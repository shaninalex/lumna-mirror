import { createFeatureSelector, createSelector } from "@ngrx/store"
import { CommentsState, commentsAdapter } from "./comment.reducer"
import { byMostRecent } from "@client/shared/common"

export const selectCommentsFeature = createFeatureSelector<CommentsState>("comments")
export const commentsSelectors = commentsAdapter.getSelectors(selectCommentsFeature)
export const selectAllComments = commentsSelectors.selectAll

export const selectComments = (entity_id: number, entity_type: string) =>
	createSelector(selectAllComments, comments =>
		comments.filter(c => c.entity_id === entity_id && c.entity_type === entity_type).sort(byMostRecent)
	)
