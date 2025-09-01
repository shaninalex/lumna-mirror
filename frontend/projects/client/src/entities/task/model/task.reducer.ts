import {createEntityAdapter, EntityAdapter, EntityState} from '@ngrx/entity';
import {createReducer, on} from '@ngrx/store';
import {Task} from './task.model'
import {SetTasksActions} from './task.actions';

export interface TasksState extends EntityState<Task> {}
export const tasksAdapter: EntityAdapter<Task> = createEntityAdapter<Task>();
export const tasksReducer = createReducer(
    tasksAdapter.getInitialState(),
    on(SetTasksActions, (state, action) => tasksAdapter.addMany(action.payload, state)),
)
