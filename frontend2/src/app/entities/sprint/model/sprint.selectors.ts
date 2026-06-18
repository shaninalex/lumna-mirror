import { createFeatureSelector, createSelector } from "@ngrx/store";
import { sprintAdapter, SprintState } from "./sprint.store";
import { SprintModel } from "./sprint.model";

const feature = createFeatureSelector<SprintState>("sprint");
const selectors = sprintAdapter.getSelectors();

export const selectSprintList = createSelector(feature, (state) =>
    selectors.selectAll(state)
);

export const selectSprint = (id: number) =>
    createSelector(selectSprintList, (list) =>
        list.filter((a: SprintModel) => a.id === id)
    );

export const selectSprintByProject = (projectId: number) =>
    createSelector(selectSprintList, (list) =>
        list.filter((a: SprintModel) => a.project_id === projectId)
    );
