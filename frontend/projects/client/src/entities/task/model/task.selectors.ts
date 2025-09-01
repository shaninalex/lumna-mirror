import {createFeatureSelector, createSelector} from '@ngrx/store';
import {tasksAdapter, TasksState} from '@client/entities/task';

export const selectTasksFeature = createFeatureSelector<TasksState>('task');
export const tasksSelectors = tasksAdapter.getSelectors();

// select only tasks from the single project instead of all
export const selectTasks = (projectID: string) => createSelector(
    selectTasksFeature,
    tasksSelectors.selectAll,
);
