import { createFeatureSelector, createSelector } from "@ngrx/store";
import type { TaskState } from "./task.store";
import { taskAdapter } from "./task.store";
import type { TaskModel } from "./task.model";

const feature = createFeatureSelector<TaskState>("task");
const selectors = taskAdapter.getSelectors();
export const selectTaskList = createSelector(feature, (state) => selectors.selectAll(state));
export const selectTask = (id: number) => createSelector(selectTaskList, (list) => list.filter((a: TaskModel) => a.id === id));
export const selectTasksByProject = (projectId: number) => createSelector(
    selectTaskList,
    (list) => (!projectId ? [] : list.filter((a) => a.project_id === projectId))
);
