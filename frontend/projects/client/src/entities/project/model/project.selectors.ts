import {createFeatureSelector, createSelector} from '@ngrx/store';
import {projectsAdapter, ProjectState} from '@client/entities/project';

export const selectProjectsFeature = createFeatureSelector<ProjectState>('project');
export const projectsSelectors = projectsAdapter.getSelectors();
export const selectProjects = createSelector(
    selectProjectsFeature,
    projectsSelectors.selectAll,
);
