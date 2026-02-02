import { UserModel } from './user.model';
import { createReducer, on } from '@ngrx/store';
import { actionUserClear, actionUserSet } from './user.actions';

export interface UserState {
    user: UserModel | null;
}

const initialState: UserState = {
    user: null,
};

export const userReducer = createReducer(
    initialState,
    on(actionUserSet, (state, action) => ({ user: action.user })),
    on(actionUserClear, (state, action) => ({ user: null })),
);
