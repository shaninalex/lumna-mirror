import {sessionReducer} from '@client/entities/auth';
import {projectsReducer} from '@client/entities/project';
import {tasksReducer} from '@client/entities/task';

import * as sessionEffects from '@client/entities/auth/model/session.effects';
import * as projectEffects from '@client/entities/project/model/project.effects';
import * as taskEffects from '@client/entities/task/model/task.effects';
import * as appEffects from '@client/shared/store/app.effects'

export const reducers = {
    session: sessionReducer,
    project: projectsReducer,
    task: tasksReducer,
}

export const effects = [
    appEffects,
    sessionEffects,
    projectEffects,
    taskEffects,
]
