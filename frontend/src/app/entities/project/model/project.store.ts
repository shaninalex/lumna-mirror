import { createReducer, on } from '@ngrx/store';
import { ProjectModel } from './project.model';
import { createEntityAdapter, EntityState } from '@ngrx/entity';
import * as ProjectActions from './project.actions';

export interface ProjectState extends EntityState<ProjectModel> {}
export const projectsAdapter = createEntityAdapter<ProjectModel>();
export const initialState = projectsAdapter.getInitialState();

export const projectsReducer = createReducer(
    initialState,
    on(ProjectActions.actionProjectsSetList, (state, action) =>
        projectsAdapter.setAll(action.projects, state),
    ),
    on(ProjectActions.actionProjectsAdd, (state, action) =>
        projectsAdapter.addOne(action.project, state),
    ),
);
