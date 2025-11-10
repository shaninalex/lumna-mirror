import { createEntityAdapter, EntityAdapter, EntityState } from "@ngrx/entity"
import { createReducer, on } from "@ngrx/store"
import * as _ from "lodash"
import { Task } from "./task.model"
import {
	TaskCreateCommentSuccessAction,
	TaskDeleteSuccessAction,
	TaskListSetActions,
	TaskSetAction,
	TaskUpdateAction,
} from "./task.actions"

export interface TasksState extends EntityState<Task> {}
export const tasksAdapter: EntityAdapter<Task> = createEntityAdapter<Task>()
export const tasksReducer = createReducer(
	tasksAdapter.getInitialState(),
	on(TaskListSetActions, (state, action) => tasksAdapter.addMany(action.payload, state)),
	on(TaskSetAction, (state, action) => tasksAdapter.addOne(action.payload, state)),
	on(TaskUpdateAction, (state, action) =>
		tasksAdapter.updateOne(
			{
				id: action.payload.id,
				changes: action.payload,
			},
			state
		)
	),
	on(TaskDeleteSuccessAction, (state, action) => tasksAdapter.removeOne(action.taskId, state)),
	on(TaskCreateCommentSuccessAction, (state, action) => {
		const task = _.cloneDeep(state.entities[action.payload.task_id])
		if (!task) return state
		const comments = [...task.comments]
		comments.push(action.payload)
		return tasksAdapter.updateOne(
			{
				id: action.payload.task_id,
				changes: {
					comments: comments,
				},
			},
			state
		)
	})
)
