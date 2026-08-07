import { routerReducer } from '@ngrx/router-store';

import { UserEffects, userReducer } from '@entities/user';

import { SessionEffects } from '@core/store/session.effects';
import { sessionReducer } from '@core/store/session.store';

export const effects = [
    SessionEffects,
    UserEffects,
];

export const reducers = {
    user: userReducer,
    session: sessionReducer,
    
    router: routerReducer,
};
