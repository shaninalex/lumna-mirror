import { createFeatureSelector, createSelector } from '@ngrx/store';
import type { WorkspaceState } from './workspace.store';
import { workspaceAdapter } from './workspace.store';
import type { WorkspaceModel } from './workspace.model';

const feature = createFeatureSelector<WorkspaceState>('workspace');
const entitySelectors = workspaceAdapter.getSelectors();

const currentWorkspaceId = createSelector(feature, (state) => state?.currentId ?? null);

export const selectWorkspaces = {
    all: createSelector(feature, entitySelectors.selectAll),
    entities: createSelector(feature, entitySelectors.selectEntities),
    total: createSelector(feature, entitySelectors.selectTotal),
    byId: (id: number) =>
        createSelector(entitySelectors.selectAll, (list) =>
            list.find((a: WorkspaceModel) => a.id === id),
        ),
    currentWorkspaceId: currentWorkspaceId,
    currentWorkspace: createSelector(
        entitySelectors.selectAll,
        currentWorkspaceId,
        (list, currentId) => list.find((w) => w.id === currentId) ?? null,
    ),
};
