import { createAction, props } from "@ngrx/store";
import { UserModel } from "@entities/user";
import { Error } from "@shared/models";

export const actionSessionAuthenticateStart = createAction(
    "[Session] start authenticate",
    props<{ email: string; password: string }>()
);

export const actionSessionAuthenticatedSuccessfull = createAction(
    "[Session] authenticated successfull",
    props<{ user: UserModel }>()
);

export const actionLoginFailed = createAction(
    "[Session] login failed",
    props<{ errors: Error[] }>()
);

export const actionSessionAuthenticated = createAction(
    "[Session] authenticated"
);
export const actionSessionLoggingOut = createAction("[Session] logging out");
export const actionSessionLoggedOut = createAction("[Session] logged out");
export const actionSessionFailed = createAction("[Session] failed");
