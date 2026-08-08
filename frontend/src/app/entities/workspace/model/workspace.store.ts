import { createEntityAdapter, EntityState } from "@ngrx/entity";
import { createReducer, on } from "@ngrx/store";
import { actionWorkspaceSetCurrent, actionWorkspaceSetList } from "./workspace.actions";
import { WorkspaceModel } from "./workspace.model";

export interface WorkspaceState extends EntityState<WorkspaceModel> {
    currentId: number | null;
}
export const workspaceAdapter = createEntityAdapter<WorkspaceModel>();
const initialState: WorkspaceState = workspaceAdapter.getInitialState({
    currentId: null
});

export const workspaceReducer = createReducer(
    initialState,
    on(actionWorkspaceSetList, (state, { list }) =>
        workspaceAdapter.setAll(list, state)
    ),
    on(actionWorkspaceSetCurrent, (state, { id }) => ({
        ...state,
        currentId: id
    }))
);

