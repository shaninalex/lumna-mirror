import { createFeatureSelector, createSelector } from '@ngrx/store';
import { projectsAdapter, ProjectState } from './project.store';

export const selectProjectsFeature = createFeatureSelector<ProjectState>('project');
export const projectsSelectors = projectsAdapter.getSelectors();

export const selectProjects = createSelector(selectProjectsFeature, (state) =>
    projectsSelectors.selectAll(state),
);

export const selectProjectByID = (id: number) =>
    createSelector(selectProjectsFeature, (state: ProjectState) =>
        projectsSelectors.selectAll(state).find((p) => p.id === id),
    );

export const selectProjectsByWorkspaceID = (workspaceId: number) =>
    createSelector(selectProjectsFeature, (state: ProjectState) =>
        projectsSelectors.selectAll(state).filter((p) => p.workspace_id === workspaceId),
    );
