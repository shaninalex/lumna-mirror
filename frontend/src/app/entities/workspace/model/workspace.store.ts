import { createEntityAdapter, EntityState } from "@ngrx/entity";
import { createReducer, on } from "@ngrx/store";
import { actionWorkspaceSetList } from "./workspace.actions";
import { WorkspaceModel } from "./workspace.model";

export interface WorkspaceState extends EntityState<WorkspaceModel> {}
export const workspaceAdapter = createEntityAdapter<WorkspaceModel>();
const initialState = workspaceAdapter.getInitialState();

export const workspaceReducer = createReducer(
    initialState,
    on(actionWorkspaceSetList, (state, { list }) =>
        workspaceAdapter.addMany(list, state)
    )
);
