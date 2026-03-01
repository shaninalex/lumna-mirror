import { createFeatureSelector, createSelector } from '@ngrx/store';
import { activityAdapter, ActivityState } from './activity.store';
import {ActivityModel} from '@entities/activity/model/activity.model';

const selectTaskFeature = createFeatureSelector<ActivityState>('activity');
const activitySelectors = activityAdapter.getSelectors();
const selectTasks = createSelector(selectTaskFeature, (state) => activitySelectors.selectAll(state));

export const selectActivity = (entity_id: number, entity_type: string) =>
    createSelector(selectTasks, (list) =>
        list.filter((a: ActivityModel) => a.entity_id === entity_id && a.entity_type === entity_type),
    );


