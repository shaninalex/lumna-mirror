import { eventGroup } from '@ngrx/signals/events';
import { type } from '@ngrx/signals';
import { TaskModel, TaskPayloadModel } from '@entities/task';
import { KanbanBoardChangeOrderPayload } from '@features/kanban-board/model/kanban.model';

export const taskEvents = eventGroup({
    source: 'Task',
    events: {
        getTasks: type<{ board_id: string }>(),
        create: type<{ board_id: string; data: TaskPayloadModel }>(),
        patch: type<{ task_id: string; data: TaskPayloadModel }>(),
        changeOrder: type<KanbanBoardChangeOrderPayload>(),

        delete: type<{ task_id: string }>(),

        setTasks: type<TaskModel[]>(),
        setTask: type<TaskModel>(),

        failed: type<any>(),

        _patchSuccess: type<TaskModel>(),
        _deleteSuccess: type<{ task_id: string }>(),
    },
});
