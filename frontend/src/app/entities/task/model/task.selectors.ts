import { createFeatureSelector, createSelector } from '@ngrx/store';
import { taskAdapter, TaskState } from './task.store';

const selectTaskFeature = createFeatureSelector<TaskState>('task');
const taskSelectors = taskAdapter.getSelectors();
const selectTasks = createSelector(selectTaskFeature, (state) => taskSelectors.selectAll(state));

export const selectTasksByColumns = (columns_id: number[]) =>
    createSelector(selectTasks, (tasks) =>
        tasks.filter((task) => columns_id.includes(task.column_id)),
    );
