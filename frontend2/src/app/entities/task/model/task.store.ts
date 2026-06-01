import { createEntityAdapter, EntityState } from "@ngrx/entity";
import { createReducer, on } from "@ngrx/store";
import { actionTaskSetList } from "./task.actions";
import { TaskModel } from "./task.model";

export interface TaskState extends EntityState<TaskModel> {}
export const taskAdapter = createEntityAdapter<TaskModel>();
const initialState = taskAdapter.getInitialState();

export const taskReducer = createReducer(
    initialState,
    on(actionTaskSetList, (state, { tasks }) =>
        taskAdapter.addMany(tasks, state)
    )
);
