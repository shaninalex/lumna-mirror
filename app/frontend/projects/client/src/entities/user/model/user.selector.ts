import {createFeatureSelector, createSelector} from '@ngrx/store';
import {UserState} from '@client/entities/user';

const selectUserFeature = createFeatureSelector<UserState>('user');
export const selectUser = createSelector(
    selectUserFeature,
    (state: UserState) => state.user
);
export const selectUserState = createSelector(
    selectUserFeature,
    (state: UserState) => state
);

