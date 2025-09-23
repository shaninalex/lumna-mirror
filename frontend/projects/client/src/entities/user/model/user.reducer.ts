import {createReducer, on} from '@ngrx/store';
import {UserModel} from '@client/entities/user';
import {SetUserAction} from '@client/entities/user/model/user.actions';

export interface UserState {
    user: UserModel | null
    loaded: boolean
}

const initialState: UserState = {
    user: null,
    loaded: false,
}

export const userReducer = createReducer(
    initialState,
    on(SetUserAction, (state, action) => ({...state, user: action.payload, loaded: true}))
)
