import { createFeatureSelector, createSelector } from '@ngrx/store';
import type { ColumnState } from './column.store';
import { statusAdapter } from './column.store';

const feature = createFeatureSelector<ColumnState>('column');
const entitySelectors = statusAdapter.getSelectors();

export const selectColumns = {
    all: createSelector(feature, entitySelectors.selectAll),
    entities: createSelector(feature, entitySelectors.selectEntities),
    total: createSelector(feature, entitySelectors.selectTotal),
    byId: (id: number) =>
        createSelector(feature, (state) =>
            entitySelectors.selectAll(state).find((a) => a.id === id),
        ),
    byListId: (listId: number) =>
        createSelector(feature, (state) =>
            entitySelectors.selectAll(state).filter((a) => a.board_id === listId),
        ),
    loading: createSelector(feature, (state) => state.loading),
    error: createSelector(feature, (state) => state.errors),
};
