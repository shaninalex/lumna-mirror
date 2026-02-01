import { userReducer, UserEffects } from '@entities/user';
import { ProjectEffects, projectReducer } from '@entities/project';
import { BoardsEffects, boardReducer } from '@entities/board';
import { ColumnEffects, columnReducer } from '@entities/column';

import { SessionEffects } from '@core/store/session.effects';
import { sessionReducer } from '@core/store/session.store';
import { TaskEffects, taskReducer } from '@entities/task';
import { KanbanEffects } from '@features/kanban-board';

export const effects = [
    // Entity effects
    ProjectEffects,
    SessionEffects,
    BoardsEffects,
    ColumnEffects,
    UserEffects,
    TaskEffects,

    // Feature effects
    KanbanEffects,
];

export const reducers = {
    project: projectReducer,
    user: userReducer,
    session: sessionReducer,
    board: boardReducer,
    column: columnReducer,
    task: taskReducer,
};
