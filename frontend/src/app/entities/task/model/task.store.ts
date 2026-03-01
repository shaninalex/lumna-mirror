import { TaskModel } from './task.model';
import { createEntityAdapter, EntityState, Update } from '@ngrx/entity';
import { createReducer, on } from '@ngrx/store';
import {
    actionTaskChange,
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
    on(actionTaskSetTasks, (state, { tasks }) => taskAdapter.addMany(tasks, state)),
    on(actionTaskUpsert, (state, { task }) => taskAdapter.upsertOne(task, state)),
    on(actionTaskDeleteSuccess, (state, { taskId }) => taskAdapter.removeOne(taskId, state)),

    // Change order for multiple tasks
    on(actionTaskChangeOrder, (state, action) => {
        const { columnId, tasks, columns } = action;

        // CASE 1: reorder tasks inside one list
        if (columnId && tasks) {
            const updates: Update<TaskModel>[] = tasks.map((t) => ({
                id: t.id,
                changes: {
                    order: t.order,
                },
            }));
            return taskAdapter.updateMany(updates, state);
        }

        // CASE 2: move tasks between columns
        if (columns) {
            const updates: Update<TaskModel>[] = columns.flatMap((column) =>
                (column.tasks ?? []).map((task) => ({
                    id: task.id,
                    changes: {
                        order: task.order,
                        column_id: column.id,
                    },
                })),
            );

            return taskAdapter.updateMany(updates, state);
        }

        return state;
    }),

    on(actionTaskChange, (state, action) => taskAdapter.updateOne({
        id: action.task_id,
        changes: action.data,
    }, state)),
);
