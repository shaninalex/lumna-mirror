import { userReducer, UserEffects } from '@entities/user';
import { ProjectEffects, projectReducer } from '@entities/project';
import { BoardsEffects, boardReducer } from '@entities/board';
import { ColumnEffects, columnReducer } from '@entities/column';

import { SessionEffects } from '@core/store/session.effects';
import { sessionReducer } from '@core/store/session.store';

export const effects = [ProjectEffects, SessionEffects, BoardsEffects, ColumnEffects, UserEffects];

export const reducers = {
    project: projectReducer,
    user: userReducer,
    session: sessionReducer,
    board: boardReducer,
    column: columnReducer,
};
