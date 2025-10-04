import {createEntityAdapter, EntityAdapter, EntityState} from '@ngrx/entity';
import {createReducer, on} from '@ngrx/store';
import {Task} from './task.model'
import {ChangeTaskStatusSuccessAction, SetTaskAction, SetTaskListActions} from './task.actions';

export interface TasksState extends EntityState<Task> {}
export const tasksAdapter: EntityAdapter<Task> = createEntityAdapter<Task>();
export const tasksReducer = createReducer(
    tasksAdapter.getInitialState(),
    on(SetTaskListActions, (state, action) => tasksAdapter.addMany(action.payload, state)),
    on(SetTaskAction, (state, action) => tasksAdapter.addOne(action.payload, state)),
    on(ChangeTaskStatusSuccessAction, (state, action) => tasksAdapter.updateOne({
        id: action.payload.id,
        changes: action.payload,
    }, state))
)
