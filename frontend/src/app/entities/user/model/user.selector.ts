import { createFeatureSelector, createSelector } from '@ngrx/store';
import { UserState } from './user.store';

const selectUserFeature = createFeatureSelector<UserState>('user');
export const selectUser = createSelector(selectUserFeature, (state: UserState) => state.user);
export const selectUserState = createSelector(selectUserFeature, (state: UserState) => state);
