import {sessionReducer} from '@client/entities//session';

export const reducers = {
    session: sessionReducer,
}

import * as sessionEffects from '@client/entities/session/model/session.effects';

export const effects = [
    sessionEffects,
]
