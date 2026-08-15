import { createEntityAdapter, EntityState } from "@ngrx/entity";
import { createReducer, on } from "@ngrx/store";
import { actionTaskCreateSuccess, actionTaskSetList } from "./task.actions";
import { TaskModel } from "./task.model";

export interface TaskState extends EntityState<TaskModel> {}
export const taskAdapter = createEntityAdapter<TaskModel>({
    sortComparer: (a, b) => new Date(b.created_at).getTime() - new Date(a.created_at).getTime(),
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
