import { createReducer, on } from '@ngrx/store';
import { ProjectModel } from './project.model';
import { createEntityAdapter, EntityState } from '@ngrx/entity';
import {
    actionProjectDeleteSuccess,
    actionProjectSetCurrent,
    actionProjectsSetList,
    actionProjectUpsert,
} from '@entities/project';
import { actionWorkspaceSetCurrent } from '@entities/workspace';

export interface ProjectState extends EntityState<ProjectModel> {
    currentId: number | null;
}
export const projectsAdapter = createEntityAdapter<ProjectModel>();
export const initialState: ProjectState = projectsAdapter.getInitialState({
    currentId: null,
});

export const projectReducer = createReducer(
    initialState,
    on(actionProjectsSetList, (state, action) =>
        projectsAdapter.addMany(action.projects, state),
    ),
    on(actionProjectUpsert, (state, action) =>
        projectsAdapter.upsertOne(action.project, state),
    ),
    on(actionProjectDeleteSuccess, (state, action) =>
        projectsAdapter.removeOne(action.project_id, {
            ...state,
            currentId: state.currentId === action.project_id ? null : state.currentId,
        })
    ),
    on(actionProjectSetCurrent, (state, { id }) => ({ ...state, currentId: id })),
    // Drop the current project as soon as it no longer belongs to the active workspace
    on(actionWorkspaceSetCurrent, (state, { id }) => {
        if (state.currentId === null) {
            return state;
        }
        const current = state.entities[state.currentId];
        return current && current.workspace_id === id
            ? state
            : { ...state, currentId: null };
    })
);
