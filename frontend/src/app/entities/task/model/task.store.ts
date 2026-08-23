import type { EntityState } from "@ngrx/entity";
import { createEntityAdapter } from "@ngrx/entity";
import { createReducer, on } from "@ngrx/store";
import { actionTaskCreateSuccess, actionTaskSetList } from "./task.actions";
import type { TaskModel } from "./task.model";

export type TaskState = EntityState<TaskModel>
export const taskAdapter = createEntityAdapter<TaskModel>({
    sortComparer: (a, b) => a.order - b.order,
});
const initialState = taskAdapter.getInitialState();

export const taskReducer = createReducer(
    initialState,
    on(actionTaskSetList, (state, { tasks }) =>
        taskAdapter.addMany(tasks, state)
    ),
    on(actionTaskCreateSuccess, (state, { task }) =>
        taskAdapter.addOne(task, state)
    )
);
