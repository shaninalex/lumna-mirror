import { actionUser } from './user.actions';
import type { UserModel } from './user.model';
import { createFeature, createReducer, on } from '@ngrx/store';

export interface UserState {
    user: UserModel | null;
}

const initialState: UserState = {
    user: null,
};

export const userReducer = createReducer(
    initialState,
    on(actionUser.set, (state, action) => ({ user: action.user })),
    on(actionUser.clear, () => ({ user: null })),
);

export const userFeature = createFeature({
    name: 'user',
    reducer: userReducer,
});
