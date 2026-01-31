import { userReducer } from '@entities/user';
import { ProjectsEffects, projectsReducer } from '@entities/project';
import { BoardsEffects, boardsReducer } from '@entities/board';
import { ColumnEffects, columnReducer } from '@entities/column';

import { SessionEffects } from '@core/store/session.effects';
import { sessionReducer } from '@core/store/session.store';

export const effects = [ProjectsEffects, SessionEffects, BoardsEffects, ColumnEffects];

export const reducers = {
    project: projectsReducer,
    user: userReducer,
    session: sessionReducer,
    board: boardsReducer,
    column: columnReducer,
};
