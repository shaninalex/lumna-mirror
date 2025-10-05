import {createAction, props} from '@ngrx/store';
import {CreateTaskInput, Task} from './task.model';
import {ChangeStatusPayload} from '@client/features/project/board-view-feature/api';

export const GetTasksActions = createAction(
    "[task] get tasks",
    props<{ projectId: number }>(),
)

export const SetTaskListActions = createAction(
    "[task] set tasks",
    props<{ payload: Task[] }>(),
)

export const SetTaskAction = createAction(
    "[task] set task",
    props<{ payload: Task }>(),
)

export const CreateTaskAction = createAction(
    "[task] create task",
    props<{ projectId: number, payload: CreateTaskInput }>(),
)

export const ChangeTaskStatusAction = createAction(
    "[task] change status",
    props<{taskId: number, payload: ChangeStatusPayload}>(),
)

export const ChangeTaskStatusSuccessAction = createAction(
    "[task] change status",
    props<{payload: Task}>(),
)
