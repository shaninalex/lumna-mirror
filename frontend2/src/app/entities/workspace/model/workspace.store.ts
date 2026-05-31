import { createEntityAdapter, EntityState } from "@ngrx/entity";
import { createReducer, on } from "@ngrx/store";
import {
    actionWorkspaceCreateSuccess,
    actionWorkspaceSetCurrent,
    actionWorkspaceSetList
} from "./workspace.actions";
import { WorkspaceModel } from "./workspace.model";

export interface WorkspaceState extends EntityState<WorkspaceModel> {
    loaded: boolean;
    loading: boolean;
    currentWorkspaceId: number | null;
}
export const workspaceAdapter = createEntityAdapter<WorkspaceModel>();
const initialState = workspaceAdapter.getInitialState({
    loaded: false,
    loading: false,
    currentWorkspaceId: null
});

export const workspaceReducer = createReducer(
    initialState,
    on(actionWorkspaceSetList, (state, { list }) =>
        workspaceAdapter.addMany(list, {
            ...state,
            loading: false,
            loaded: true
        })
    ),
    on(actionWorkspaceCreateSuccess, (state, { data }) =>
        workspaceAdapter.addOne(data, state)
    ),
    on(actionWorkspaceSetCurrent, (state, { id }) => ({
        ...state,
        currentWorkspaceId: id
    }))
);
