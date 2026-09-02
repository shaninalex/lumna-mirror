import type { EntityState } from "@ngrx/entity";
import { createEntityAdapter } from "@ngrx/entity";
import { createReducer, on } from "@ngrx/store";
import type { TaskModel } from "./task.model";
import { actionTask } from "./task.actions";

export type TaskState = EntityState<TaskModel>
export const taskAdapter = createEntityAdapter<TaskModel>({
    // sortComparer: (a, b) => a.order - b.order,
});
const initialState = taskAdapter.getInitialState();

export const taskReducer = createReducer(
    initialState,
    on(actionTask.getListSuccess, (state, { tasks }) =>
        taskAdapter.addMany(tasks, state)
    ),
    on(actionTask.createSuccess, (state, { task }) =>
        taskAdapter.addOne(task, state)
    )
);
