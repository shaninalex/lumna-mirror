import { EntityState, EntityAdapter, createEntityAdapter } from '@ngrx/entity';
import {Project} from '@client/entities/project/model/project.model';
import {createReducer, on} from '@ngrx/store';
import {SetProjectsAction} from '@client/entities/project/model/project.actions';

export interface ProjectState extends EntityState<Project> {}
export const projectsAdapter: EntityAdapter<Project> = createEntityAdapter<Project>();
export const projectsReducer = createReducer(
    projectsAdapter.getInitialState(),
    on(SetProjectsAction, (state, action) => projectsAdapter.addMany(action.payload, state)),
)
