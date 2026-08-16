import { createAction, props } from '@ngrx/store';
import type { UserModel } from '@entities/user';

export const actionSessionAuthenticateStart = createAction(
    '[Session] start authenticate',
    props<{ email: string; password: string }>(),
);

export const actionSessionAuthenticatedSuccessfull = createAction(
    '[Session] authenticated successfull',
    props<{ user: UserModel }>(),
);

export const actionSessionAuthenticated = createAction('[Session] authenticated');
export const actionSessionLoggingOut = createAction('[Session] logging out');
export const actionSessionLoggedOut = createAction('[Session] logged out');
