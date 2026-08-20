import {
    actionListCreateSuccess,
    actionListDeleteSuccess,
    actionListPatchSuccess,
    actionListSet,
    actionListSetList,
} from './list.actions'
import type { EntityState } from '@ngrx/entity';
import { createEntityAdapter } from '@ngrx/entity';
import { createReducer, on } from '@ngrx/store';
import type { ListModel } from './list.model';

export type ListState = EntityState<ListModel>;
export const listAdapter = createEntityAdapter<ListModel>();
const initialState = listAdapter.getInitialState();

export const listReducer = createReducer(
    initialState,
    on(actionListSetList, (state, { lists }) => listAdapter.addMany(lists, state)),
    on(actionListSet, (state, { list }) => listAdapter.upsertOne(list, state)),
    on(actionListCreateSuccess, (state, { list }) => listAdapter.upsertOne(list, state)),
    on(actionListPatchSuccess, (state, { list }) => listAdapter.upsertOne(list, state)),
    on(actionListDeleteSuccess, (state, { listId }) => listAdapter.removeOne(listId, state)),
);
