import {sessionReducer} from '@client/entities//session';
import {projectsReducer} from '@client/entities/project';
import {tasksReducer} from '@client/entities/task';

import * as sessionEffects from '@client/entities/session/model/session.effects';
import * as projectEffects from '@client/entities/project/model/project.effects';
import * as taskEffects from '@client/entities/task/model/task.effects';

export const reducers = {
    session: sessionReducer,
    project: projectsReducer,
    task: tasksReducer,
}

export const effects = [
    sessionEffects,
    projectEffects,
    taskEffects,
]
