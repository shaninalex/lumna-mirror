import {createReducer, on} from '@ngrx/store';
import {UserModel} from '@client/entities/user';

export interface UserState {
    user: UserModel | undefined
}

const initialState: UserState = {
    user: undefined
}

export const userReducer = createReducer(
    initialState,
)
