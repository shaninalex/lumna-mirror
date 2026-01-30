import { createReducer, on } from '@ngrx/store';
import { ProjectModel } from './project.model';
import { createEntityAdapter, EntityAdapter, EntityState } from '@ngrx/entity';
import * as ProjectActions from './project.actions';

export interface ProjectState extends EntityState<ProjectModel> {}
export const projectsAdapter: EntityAdapter<ProjectModel> = createEntityAdapter<ProjectModel>();
export const initialState: ProjectState = projectsAdapter.getInitialState();

export const projectsReducer = createReducer(
    initialState,
    on(ProjectActions.actionProjectsSet, (state, action) =>
        projectsAdapter.setAll(action.projects, state),
    ),
);
