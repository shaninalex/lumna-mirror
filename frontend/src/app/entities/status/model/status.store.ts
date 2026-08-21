import { createReducer, on } from '@ngrx/store';
import { createEntityAdapter, type EntityState } from '@ngrx/entity';
import type { Error } from '@shared/models';
import type { StatusModel } from './status.model';
import { statusActions } from './status.actions';

export interface StatusState extends EntityState<StatusModel> {
    loading: boolean;
    errors: Error[];
}

export const statusAdapter = createEntityAdapter<StatusModel>({
    sortComparer: (a, b) => b.order - a.order,
});

const initialState: StatusState = statusAdapter.getInitialState({
    loading: false,
    errors: [],
});

export const statusReducer = createReducer(
    initialState,
    on(statusActions.createSuccess, (state, action) => statusAdapter.addOne(action.status, state)),
    on(statusActions.loadByListIdSuccess, (state, action) =>
        statusAdapter.addMany(action.statuses, state),
    ),
    on((statusActions.loadByListIdFailed, statusActions.createFailed), (state, action) => ({
        ...state,
        errors: action.errors,
    })),
);
