import {ProjectState} from '@client/entities/project';
import {TasksState} from '@client/entities/task'
import {SessionState} from '@client/entities/auth'

export interface AppState {
    session: SessionState
    project: ProjectState
    task: TasksState
}
