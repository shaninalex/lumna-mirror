import { createFeatureSelector, createSelector } from '@ngrx/store';
import { projectsAdapter, ProjectState } from './project.store';

export const selectProjectsFeature = createFeatureSelector<ProjectState>('project');
export const projectsSelectors = projectsAdapter.getSelectors();

export const selectProjects = createSelector(selectProjectsFeature, (state) =>
    projectsSelectors.selectAll(state),
);

export const selectProjectByID = (id: string) =>
    createSelector(selectProjectsFeature, (state: ProjectState) =>
        projectsSelectors.selectAll(state).find((p) => p.id === id),
    );
