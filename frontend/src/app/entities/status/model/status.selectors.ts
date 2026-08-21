import { createFeatureSelector, createSelector } from '@ngrx/store';
import type { StatusState } from './status.store';
import { statusAdapter } from './status.store';

const feature = createFeatureSelector<StatusState>('status');

const entitySelectors = statusAdapter.getSelectors();

export const statusSelectors = {
    all: createSelector(feature, entitySelectors.selectAll),

    entities: createSelector(feature, entitySelectors.selectEntities),

    total: createSelector(feature, entitySelectors.selectTotal),

    byId: (id: number) =>
        createSelector(feature, (state) =>
            entitySelectors.selectAll(state).find((a) => a.id === id),
        ),

    byListId: (listId: number) =>
        createSelector(feature, (state) =>
            entitySelectors.selectAll(state).find((a) => a.list_id === listId),
        ),

    byProjectId: (projectId: number) =>
        createSelector(feature, (state) =>
            entitySelectors.selectAll(state).find((a) => a.project_id === projectId),
        ),

    loading: createSelector(
        feature,
        state => state.loading,
    ),

    error: createSelector(
        feature,
        state => state.errors,
    ),
};
