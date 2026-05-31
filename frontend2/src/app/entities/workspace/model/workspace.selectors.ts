import { createFeatureSelector, createSelector } from "@ngrx/store";
import { workspaceAdapter, WorkspaceState } from "./workspace.store";
import { WorkspaceModel } from "./workspace.model";

export const workspacesFeature =
    createFeatureSelector<WorkspaceState>("workspace");
const selectors = workspaceAdapter.getSelectors();

export const selectWorkspaceList = createSelector(workspacesFeature, (state) =>
    selectors.selectAll(state)
);

export const selectWorkspace = (id: number) =>
    createSelector(selectWorkspaceList, (list) =>
        list.filter((a: WorkspaceModel) => a.id === id)
    );

export const selectWorkspaceById = (id: number) =>
    createSelector(selectWorkspaceList, (list) =>
        list.find((a: WorkspaceModel) => a.id === id)
    );

export const selectWorkspaceListLoaded = createSelector(
    workspacesFeature,
    (s) => s.loaded
);

export const selectWorkspaceCurrentWorkspaceId = createSelector(
    workspacesFeature,
    (s) => s.currentWorkspaceId
);

export const selectWorkspaceCurrent = createSelector(
    selectWorkspaceCurrentWorkspaceId,
    selectWorkspaceList,
    (id, workspaces) => {
        if (!id) {
            return undefined;
        }
        return workspaces.find((workspace) => workspace.id === id);
    }
);
