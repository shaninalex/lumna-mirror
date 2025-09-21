import {ProjectState} from '@client/entities/project';
import {TasksState} from '@client/entities/task'
import {UserState} from '@client/entities/user';

export interface AppState {
    project: ProjectState
    task: TasksState
    user: UserState
}
