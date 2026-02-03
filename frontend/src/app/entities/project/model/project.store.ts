import { createReducer, on } from '@ngrx/store';
import { ProjectModel } from './project.model';
import { createEntityAdapter, EntityState } from '@ngrx/entity';
import { actionProjectDeleteSuccess, actionProjectsSetList, actionProjectUpsert } from '@entities/project';

export interface ProjectState extends EntityState<ProjectModel> {}
export const projectsAdapter = createEntityAdapter<ProjectModel>();
export const initialState = projectsAdapter.getInitialState();

export const projectReducer = createReducer(
    initialState,
    on(actionProjectsSetList, (state, action) =>
        projectsAdapter.addMany(action.projects, state),
    ),
    on(actionProjectUpsert, (state, action) =>
        projectsAdapter.upsertOne(action.project, state),
    ),
    on(actionProjectDeleteSuccess, (state, action) =>
        projectsAdapter.removeOne(action.project_id, state)
    )
);
