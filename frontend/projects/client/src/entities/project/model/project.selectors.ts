import {createFeatureSelector, createSelector} from '@ngrx/store';
import {byMostRecent} from '@client/shared/common';
import {Project, projectsAdapter, ProjectState} from '@client/entities/project';

export const selectProjectsFeature = createFeatureSelector<ProjectState>('project');
export const projectsSelectors = projectsAdapter.getSelectors();
export const selectProjects = createSelector(
    selectProjectsFeature,
    state => projectsSelectors.selectAll(state).sort(byMostRecent)
);

export const selectProject = (code: string) => createSelector(
    selectProjectsFeature,
    (state: ProjectState) => projectsSelectors.selectAll(state).find(p => p.project_key === code)
)
