import { projectsReducer } from '@entities/project';
import { userReducer } from '@entities/user';
import * as projectEffects from '@entities/project/model/project.effects';
import { SessionEffects } from '@core/store/session.effects';
import { sessionReducer } from '@core/store/session.store';

export const effects = [projectEffects, SessionEffects];

export const reducers = {
    project: projectsReducer,
    user: userReducer,
    session: sessionReducer,
};
