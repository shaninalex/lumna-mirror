import { createEntityAdapter, EntityState } from "@ngrx/entity";
import { createReducer, on } from "@ngrx/store";
import { actionListSetList } from "./list.actions";
import { ListModel } from "./list.model";

export interface ListState extends EntityState<ListModel> {}
export const listAdapter = createEntityAdapter<ListModel>();
const initialState = listAdapter.getInitialState();

export const listReducer = createReducer(
    initialState,
    on(actionListSetList, (state, { list }) => listAdapter.addMany(list, state))
);
