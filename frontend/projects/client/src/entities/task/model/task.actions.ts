import {createAction, props} from '@ngrx/store';
import {CreateTaskInput, Task} from './task.model';
import {ChangeStatusPayload} from '@client/features/project/board-view-feature/api';

export const TaskListGetActions = createAction(
    "[task] get tasks",
    props<{ projectId: number }>(),
)

export const TaskListSetActions = createAction(
    "[task] set tasks",
    props<{ payload: Task[] }>(),
)

export const TaskSetAction = createAction(
    "[task] set task",
    props<{ payload: Task }>(),
)

export const TaskCreateAction = createAction(
    "[task] create task",
    props<{ projectId: number, payload: CreateTaskInput }>(),
)

export const TaskChangeStatusAction = createAction(
    "[task] change status",
    props<{taskId: number, payload: ChangeStatusPayload}>(),
)

export const TaskChangeStatusSuccessAction = createAction(
    "[task] change status",
    props<{payload: Task}>(),
)
