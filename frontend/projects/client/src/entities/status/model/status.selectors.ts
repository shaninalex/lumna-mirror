import {createFeatureSelector, createSelector} from '@ngrx/store';
import {statusAdapter, StatusState} from '@client/entities/status';

export const selectStatusFeature = createFeatureSelector<StatusState>('status');
export const statusSelectors = statusAdapter.getSelectors();

export const selectStatuses = createSelector(
    selectStatusFeature,
    state => statusSelectors.selectAll(state)
);

export const selectStatus = (id: number) => createSelector(
    selectStatusFeature,
    (state: StatusState) => statusSelectors.selectAll(state).find(p => p.id === id)
)

export const selectProjectStatusList = (project_id: number) => createSelector(
    selectStatusFeature,
    (state: StatusState) => statusSelectors.selectAll(state)
        .filter(p => p.project_id === project_id)
        .sort((a, b) => a.index - b.index)
)
