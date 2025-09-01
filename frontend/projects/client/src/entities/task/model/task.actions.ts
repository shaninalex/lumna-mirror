import {createAction, props} from '@ngrx/store';
import {Task} from './task.model';

export const GetTasksActions = createAction(
    "[task] get tasks",
    props<{ projectCode: string }>(),
)


export const SetTasksActions = createAction(
    "[task] set tasks",
    props<{ payload: Task[] }>(),
)
