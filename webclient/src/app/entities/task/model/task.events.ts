import { eventGroup } from '@ngrx/signals/events';
import { type } from '@ngrx/signals';
import { TaskModel, TaskPayloadModel } from '@entities/task';

export const taskEvents = eventGroup({
    source: 'Task',
    events: {
        getTasks: type<{ board_id: number }>(),
        create: type<{ data: TaskPayloadModel }>(),
        patch: type<{ data: TaskPayloadModel }>(),
        delete: type<{ task_id: number }>(),

        setTasks: type<TaskModel[]>(),
        setTask: type<TaskModel>(),

        failed: type<any>(),

        _patchSuccess: type<TaskModel>(),
        _deleteSuccess: type<{ task_id: number }>(),
    },
});
