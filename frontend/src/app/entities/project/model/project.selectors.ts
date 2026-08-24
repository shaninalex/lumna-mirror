import { createFeatureSelector, createSelector } from '@ngrx/store';
import type { ProjectState } from './project.store';
import { projectsAdapter } from './project.store';

const feature = createFeatureSelector<ProjectState>('project');
const entitySelectors = projectsAdapter.getSelectors();

const selectCurrentProjectId = createSelector(feature, (state) => state.currentId ?? null);

const selectCurrentProject = createSelector(
    feature,
    selectCurrentProjectId,
    (state, currentId) => entitySelectors.selectAll(state).find((p) => p.id === currentId) ?? null,
);

export const selectProjects = {
    all: createSelector(feature, entitySelectors.selectAll),
    entities: createSelector(feature, entitySelectors.selectEntities),
    total: createSelector(feature, entitySelectors.selectTotal),
    byId: (id: number) =>
        createSelector(feature, (state: ProjectState) =>
            entitySelectors.selectAll(state).find((p) => p.id === id),
        ),
    byWorkspaceId: (workspaceId: number) =>
        createSelector(feature, (state: ProjectState) =>
            entitySelectors.selectAll(state).filter((p) => p.workspace_id === workspaceId),
        ),
    currentProjectId: selectCurrentProjectId,
    currentProject: selectCurrentProject,
};
