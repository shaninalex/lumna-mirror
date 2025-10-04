import {createEntityAdapter, EntityAdapter, EntityState} from '@ngrx/entity';
import {createReducer, on} from '@ngrx/store';
import {
    DeleteStatusSuccessAction, PatchStatusSuccessAction,
    SetStatusAction,
    SetStatusListActions,
    Status
} from '@client/entities/status';

export interface StatusState extends EntityState<Status> {
}

export const statusAdapter: EntityAdapter<Status> = createEntityAdapter<Status>();
export const statusReducer = createReducer(
    statusAdapter.getInitialState(),
    on(SetStatusListActions, (state, action) => statusAdapter.addMany(action.payload, state)),
    on(SetStatusAction, (state, action) => statusAdapter.addOne(action.payload, state)),
    on(PatchStatusSuccessAction, (state, action) => statusAdapter.updateOne({
        id: action.payload.id,
        changes: action.payload
    }, state)),
    on(DeleteStatusSuccessAction, (state, action) => statusAdapter.removeOne(action.statusId, state)),
)
