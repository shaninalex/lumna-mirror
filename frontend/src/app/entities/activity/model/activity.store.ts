import { ActivityModel } from './activity.model';
import { createEntityAdapter, EntityState } from '@ngrx/entity';
import { createReducer, on } from '@ngrx/store';
import {
    actionActivitySetList,
} from './activity.actions';

export interface ActivityState extends EntityState<ActivityModel> {}
export const activityAdapter = createEntityAdapter<ActivityModel>();
const initialState = activityAdapter.getInitialState();

export const activityReducer = createReducer(
    initialState,
    on(actionActivitySetList, (state, { list }) => activityAdapter.addMany(list, state)),
);
