import { createFeatureSelector, createSelector } from "@ngrx/store";
import type { WorkspaceState } from "./workspace.store";
import { workspaceAdapter } from "./workspace.store";
import type { WorkspaceModel } from "./workspace.model";

const feature = createFeatureSelector<WorkspaceState>("workspace");
const selectors = workspaceAdapter.getSelectors();
export const selectWorkspaceList = createSelector(feature, (state) =>
    selectors.selectAll(state)
);

export const selectWorkspace = (id: number) =>
    createSelector(selectWorkspaceList, (list) =>
        list.find((a: WorkspaceModel) => a.id === id)
    );

export const selectCurrentWorkspaceId = createSelector(
    feature,
    (state) => state?.currentId ?? null
);

export const selectCurrentWorkspace = createSelector(
    selectWorkspaceList,
    selectCurrentWorkspaceId,
    (list, currentId) => list.find((w) => w.id === currentId) ?? null
);

