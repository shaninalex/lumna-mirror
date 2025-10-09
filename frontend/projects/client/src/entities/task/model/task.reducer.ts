import {createEntityAdapter, EntityAdapter, EntityState} from '@ngrx/entity';
import {createReducer, on} from '@ngrx/store';
import {Task} from './task.model'
import {TaskChangeStatusSuccessAction, TaskSetAction, TaskListSetActions} from './task.actions';

export interface TasksState extends EntityState<Task> {}
export const tasksAdapter: EntityAdapter<Task> = createEntityAdapter<Task>();
export const tasksReducer = createReducer(
    tasksAdapter.getInitialState(),
    on(TaskListSetActions, (state, action) => tasksAdapter.addMany(action.payload, state)),
    on(TaskSetAction, (state, action) => tasksAdapter.addOne(action.payload, state)),
    on(TaskChangeStatusSuccessAction, (state, action) => tasksAdapter.updateOne({
        id: action.payload.id,
        changes: action.payload,
    }, state))
)
