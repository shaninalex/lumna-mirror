import { createFeatureSelector, createSelector } from "@ngrx/store";
import { workspaceAdapter, WorkspaceState } from "./workspace.store";
import { WorkspaceModel } from "./workspace.model";

const feature = createFeatureSelector<WorkspaceState>("workspace");
const selectors = workspaceAdapter.getSelectors();
export const selectWorkspaceList = createSelector(feature, (state) =>
    selectors.selectAll(state)
);

export const selectWorkspace = (id: number) =>
    createSelector(selectWorkspaceList, (list) =>
        list.filter((a: WorkspaceModel) => a.id === id)
    );
