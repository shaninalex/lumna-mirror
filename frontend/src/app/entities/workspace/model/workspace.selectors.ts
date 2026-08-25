import { createFeatureSelector, createSelector } from '@ngrx/store';
import type { WorkspaceState } from './workspace.store';
import { workspaceAdapter } from './workspace.store';

const feature = createFeatureSelector<WorkspaceState>('workspace');
const entitySelectors = workspaceAdapter.getSelectors();

const selectAll = createSelector(feature, entitySelectors.selectAll);
const currentWorkspaceId = createSelector(feature, (state) => state?.currentId ?? null);

export const selectWorkspaces = {
    all: selectAll,
    entities: createSelector(feature, entitySelectors.selectEntities),
    total: createSelector(feature, entitySelectors.selectTotal),
    byId: (id: number) => createSelector(selectAll, (list) => list.find((w) => w.id === id)),
    currentWorkspaceId: currentWorkspaceId,
    currentWorkspace: createSelector(
        selectAll,
        currentWorkspaceId,
        (list, currentId) => list.find((w) => w.id === currentId) ?? null,
    ),
};
