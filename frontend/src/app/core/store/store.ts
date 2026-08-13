import { routerReducer } from '@ngrx/router-store';
import { SessionEffects } from '@core/store/session.effects';
import { sessionReducer } from '@core/store/session.store';

export const rootEffects = [
    SessionEffects,

    // UserEffects,
    // WorkspaceEffects,
    // ProjectEffects,
];

export const rootReducers = {
    session: sessionReducer,
    router: routerReducer,

    // user: userReducer,
    // workspace: workspaceReducer,
    // project: projectReducer,
};
