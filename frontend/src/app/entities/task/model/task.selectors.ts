import { createFeatureSelector, createSelector } from '@ngrx/store';
import { taskAdapter, TaskState } from './task.store';

const selectTaskFeature = createFeatureSelector<TaskState>('task');
const taskSelectors = taskAdapter.getSelectors();
const selectTasks = createSelector(selectTaskFeature, (state) => taskSelectors.selectAll(state));

export const selectTasksByBoardId = (boardId: string) =>
    createSelector(selectTasks, (tasks) => tasks.filter((task) => task.board_id === boardId));
