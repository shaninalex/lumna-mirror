import { createReducer, on } from "@ngrx/store";
import * as SessionActions from "./session.actions";

export type SessionState = {
    authenticated: boolean;
    checked: boolean;
};

const initialState: SessionState = {
    authenticated: false,
    checked: false
};

export const sessionReducer = createReducer(
    initialState,
    on(
        SessionActions.actionSessionAuthenticatedSuccessfull,
        SessionActions.actionSessionAuthenticated,
        () => ({ authenticated: true, checked: true })
    ),
    on(
        SessionActions.actionSessionLoggedOut,
        SessionActions.actionSessionFailed,
        () => ({ authenticated: false, checked: true })
    )
);
