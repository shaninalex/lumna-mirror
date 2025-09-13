import {createAction, props} from '@ngrx/store';
import {CreateTaskDto, Task} from './task.model';

export const GetTasksActions = createAction(
    "[task] get tasks",
    props<{ projectCode: string }>(),
)

export const SetTasksActions = createAction(
    "[task] set tasks",
    props<{ payload: Task[] }>(),
)

export const SetTaskAction = createAction(
    "[task] set task",
    props<{ payload: Task }>(),
)

export const CreateTaskAction = createAction(
    "[task] create task",
    props<{ payload: CreateTaskDto }>(),
)
