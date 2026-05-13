import { createReducer, on } from '@ngrx/store';
import * as SessionActions from './session.actions';

type SessionState = {
    authenticated: boolean;
};

const initialState: SessionState = {
    authenticated: false,
};

export const sessionReducer = createReducer(
    initialState,
    on(
        SessionActions.actionSessionAuthenticatedSuccessfull,
        SessionActions.actionSessionAuthenticated,
        (state, action) => ({
            authenticated: true,
        }),
    ),
    on(SessionActions.actionSessionLoggedOut, (state, action) => ({ authenticated: false })),
);
