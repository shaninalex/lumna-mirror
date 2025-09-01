import {sessionReducer} from '@client/entities//session';
import {projectsReducer} from '@client/entities/project/model/project.reducer';

import * as sessionEffects from '@client/entities/session/model/session.effects';
import * as projectEffects from '@client/entities/project/model/project.effects';

export const reducers = {
    session: sessionReducer,
    project: projectsReducer,
}

export const effects = [
    sessionEffects,
    projectEffects,
]
