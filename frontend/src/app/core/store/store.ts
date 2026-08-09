import { routerReducer } from '@ngrx/router-store';
import { UserEffects, userReducer } from '@entities/user';
import { WorkspaceEffects, workspaceReducer } from '@entities/workspace';
import { SessionEffects } from '@core/store/session.effects';
import { sessionReducer } from '@core/store/session.store';
import { AppEffects } from './app.effects';
import { ProjectEffects, projectReducer } from '@entities/project';

export const effects = [
    SessionEffects,
    AppEffects,
    UserEffects,
    WorkspaceEffects,
    ProjectEffects,
];

export const reducers = {
    session: sessionReducer,
    router: routerReducer,
    user: userReducer,
    workspace: workspaceReducer,
    project: projectReducer,
};
