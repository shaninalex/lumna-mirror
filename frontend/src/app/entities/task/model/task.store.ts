import { TaskModel } from './task.model';
import { createEntityAdapter, EntityState } from '@ngrx/entity';
import { createReducer, on } from '@ngrx/store';
import {
    actionTaskChangeOrder,
    actionTaskDeleteSuccess,
    actionTaskSetTasks,
    actionTaskUpsert,
} from './task.actions';

export interface TaskState extends EntityState<TaskModel> {}
export const taskAdapter = createEntityAdapter<TaskModel>();
const initialState = taskAdapter.getInitialState();

export const taskReducer = createReducer(
    initialState,
    on(actionTaskSetTasks, (state, { tasks }) => taskAdapter.setAll(tasks, state)),
    on(actionTaskUpsert, (state, { task }) => taskAdapter.upsertOne(task, state)),
    on(actionTaskDeleteSuccess, (state, { taskId }) => taskAdapter.removeOne(taskId, state)),

    // Change order for multiple tasks
    on(actionTaskChangeOrder, (state, action) => {
        const { columnId, tasks, columns } = action;

        // CASE 1: reorder tasks inside one list
        if (columnId && tasks) {
            return taskAdapter.updateMany(
                tasks.map((t) => ({
                    id: t.id,
                    changes: {
                        order: t.order,
                    },
                })),
                state,
            );
        }

        // CASE 2: move tasks between columns
        if (columns) {
            const updates = columns.flatMap((list) =>
                (list.tasks ?? []).map((task) => ({
                    id: task.id,
                    changes: {
                        order: task.order,
                        listId: list.id,
                    },
                })),
            );

            return taskAdapter.updateMany(updates, state);
        }

        return state;
    }),
);
