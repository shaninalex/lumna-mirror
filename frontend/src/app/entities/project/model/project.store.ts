import { createReducer, on } from '@ngrx/store';
import { ProjectModel } from './project.model';
import { createEntityAdapter, EntityState } from '@ngrx/entity';
import * as ProjectActions from './project.actions';

export interface ProjectState extends EntityState<ProjectModel> {}
export const projectsAdapter = createEntityAdapter<ProjectModel>();
export const initialState = projectsAdapter.getInitialState();

export const projectReducer = createReducer(
    initialState,
    on(ProjectActions.actionProjectsSetList, (state, action) =>
        projectsAdapter.addMany(action.projects, state),
    ),
    on(ProjectActions.actionProjectUpsert, (state, action) =>
        projectsAdapter.upsertOne(action.project, state),
    ),
);
