import { createFeatureSelector, createSelector } from '@ngrx/store';
import type { TaskState } from './task.store';
import { taskAdapter } from './task.store';
import type { TaskModel } from './task.model';

const feature = createFeatureSelector<TaskState>('task');
const entitySelectors = taskAdapter.getSelectors();

export const selectTasks = {
    all: createSelector(feature, entitySelectors.selectAll),
    entities: createSelector(feature, entitySelectors.selectEntities),
    total: createSelector(feature, entitySelectors.selectTotal),
    byId: (id: number) =>
        createSelector(entitySelectors.selectAll, (list) =>
            list.find((a: TaskModel) => a.id === id),
        ),
    byProject: (projectId: number) =>
        createSelector(entitySelectors.selectAll, (list) =>
            list.filter((a) => a.project_id === projectId),
        ),
    countByProjectId: (projectId: number) =>
        createSelector(
            entitySelectors.selectAll,
            (list) => list.filter((a) => a.project_id === projectId).length,
        ),
    byStatusId: (statusId: number) =>
        createSelector(entitySelectors.selectAll, (list) =>
            list.filter((a) => a.status_id === statusId),
        ),
};
