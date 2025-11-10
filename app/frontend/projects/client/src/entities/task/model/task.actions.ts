import { createAction, props } from "@ngrx/store"
import { Comment, CreateTaskInput, Task, TaskDetailInput } from "./task.model"
import { ChangeStatusPayload } from "@client/features/project/board-view-feature/api"

export const TaskListGetActions = createAction("[task] get tasks", props<{ projectId: number }>())

export const TaskListSetActions = createAction("[task] set tasks", props<{ payload: Task[] }>())

export const TaskSetAction = createAction("[task] set task", props<{ payload: Task }>())

export const TaskCreateAction = createAction("[task] create task", props<{ projectId: number; payload: CreateTaskInput }>())

export const TaskChangeStatusAction = createAction("[task] change status", props<{ taskId: number; payload: ChangeStatusPayload }>())

export const TaskGetDetailsAction = createAction("[task] get details", props<{ taskId: number }>())

export const TaskUpdateAction = createAction("[task] update task", props<{ payload: Task }>())

export const TaskPatchAction = createAction("[task] patch task", props<{ taskId: number; payload: TaskDetailInput }>())

export const TaskDeleteAction = createAction("[task] delete task", props<{ taskId: number }>())

export const TaskDeleteSuccessAction = createAction("[task] delete task success", props<{ taskId: number }>())

export const TaskCreateCommentAction = createAction("[task] create comment", props<{ taskId: number; message: string }>())

export const TaskCreateCommentSuccessAction = createAction("[task] create comment success", props<{ payload: Comment }>())
