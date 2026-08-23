import { createReducer, on } from '@ngrx/store';
import { createEntityAdapter, type EntityState } from '@ngrx/entity';
import type { Error } from '@shared/models';
import type { StatusModel } from './status.model';
import { actionsStatuses } from './status.actions';

export interface StatusState extends EntityState<StatusModel> {
    loading: boolean;
    errors: Error[];
}

export const statusAdapter = createEntityAdapter<StatusModel>({
    sortComparer: (a, b) => a.meta.order - b.meta.order,
});

const initialState: StatusState = statusAdapter.getInitialState({
    loading: false,
    errors: [],
});

export const statusReducer = createReducer(
    initialState,
    on(actionsStatuses.createSuccess, (state, action) => statusAdapter.addOne(action.status, state)),
    on(actionsStatuses.loadByListIdSuccess, (state, action) =>
        statusAdapter.addMany(action.statuses, state),
    ),
    on((actionsStatuses.loadByListIdFailed, actionsStatuses.createFailed), (state, action) => ({
        ...state,
        errors: action.errors,
    })),
);
