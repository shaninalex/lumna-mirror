import {createFeatureSelector, createSelector} from '@ngrx/store';
import {byMostRecent} from '@client/shared/common';
import {projectsAdapter, ProjectState} from '@client/entities/project';

export const selectProjectsFeature = createFeatureSelector<ProjectState>('project');
export const projectsSelectors = projectsAdapter.getSelectors();
export const selectProjects = createSelector(
    selectProjectsFeature,
    state => projectsSelectors.selectAll(state).sort(byMostRecent)
);
