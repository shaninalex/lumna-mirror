import { createFeatureSelector, createSelector } from '@ngrx/store';
import { listAdapter } from './list.store';
import { BoardState } from './board.store';

const feature = createFeatureSelector<BoardState>('board');
const entitySelectors = listAdapter.getSelectors();

export const selectBoard = {
    all: createSelector(feature, entitySelectors.selectAll),
    entities: createSelector(feature, entitySelectors.selectEntities),
    total: createSelector(feature, entitySelectors.selectTotal),
    byProjectId: (projectId: number) =>
        createSelector(selectBoard.all, (lists) => {
            return lists.filter((b) => b.project_id === projectId);
        }),
    byId: (listId: number) =>
        createSelector(feature, (state) => entitySelectors.selectEntities(state)[listId] ?? null),
};
