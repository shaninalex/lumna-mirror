import {createFeatureSelector, createSelector} from '@ngrx/store';
import {tasksAdapter, TasksState} from './task.reducer';

export const selectTasksFeature = createFeatureSelector<TasksState>('task');
export const tasksSelectors = tasksAdapter.getSelectors(selectTasksFeature);
export const selectAllTasks = tasksSelectors.selectAll;
export const selectTasksByProjectID = (projectID: string) =>
    createSelector(
        selectAllTasks,
        (tasks) => tasks.filter(t => t.project_id === projectID)
    );
