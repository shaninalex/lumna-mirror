import { userReducer } from '@entities/user';
import { ProjectsEffects, projectsReducer } from '@entities/project';
import { SessionEffects } from '@core/store/session.effects';
import { sessionReducer } from '@core/store/session.store';

export const effects = [ProjectsEffects, SessionEffects];

export const reducers = {
    project: projectsReducer,
    user: userReducer,
    session: sessionReducer,
};
