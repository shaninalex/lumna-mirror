import {createEntityAdapter, EntityAdapter, EntityState} from '@ngrx/entity';
import {Task, SetTasksActions} from '@client/entities/task';
import {createReducer, on} from '@ngrx/store';

export interface TasksState extends EntityState<Task> {}
export const tasksAdapter: EntityAdapter<Task> = createEntityAdapter<Task>();
export const tasksReducer = createReducer(
    tasksAdapter.getInitialState(),
    on(SetTasksActions, (state, action) => tasksAdapter.addMany(action.payload, state)),
)
