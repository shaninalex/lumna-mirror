import { routerReducer } from '@ngrx/router-store';
import { SessionEffects } from '@core/store/session.effects';
import { sessionReducer } from '@core/store/session.store';
import { AppEffects } from './app.effects';

export const rootEffects = [
    SessionEffects,
    AppEffects,

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
