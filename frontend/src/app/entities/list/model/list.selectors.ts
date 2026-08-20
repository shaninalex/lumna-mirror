import { createFeatureSelector, createSelector } from '@ngrx/store';
import type { ListState } from './list.store';
import { listAdapter } from './list.store';

export const selectListFeature = createFeatureSelector<ListState>('list');

export const listSelectors = listAdapter.getSelectors();

export const selectLists = createSelector(selectListFeature, (state) =>
    listSelectors.selectAll(state),
);

/**
 * Select lists by project id, always sorted by `order` ASC
 */
export const selectListsByProjectId = (projectId: number) =>
    createSelector(selectLists, (lists) => {
        return lists.filter((b) => b.project_id === projectId);
    });

export const selectListById = (listId: number) =>
    createSelector(
        selectListFeature,
        (state) => listSelectors.selectEntities(state)[listId] ?? null,
    );
