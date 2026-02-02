import { createAction, props } from '@ngrx/store';
import { TaskModel, TaskPayloadModel } from './task.model';
import { KanbanBoardChangeOrderPayload } from '@features/kanban-board/model';

// LIST
export const actionTaskGetTasks = createAction('[Task] get tasks', props<{ board_id: string }>());

export const actionTaskSetTasks = createAction('[Task] set tasks', props<{ tasks: TaskModel[] }>());

// CREATE
export const actionTaskCreate = createAction(
    '[Task] create',
    props<{ board_id: string; data: TaskPayloadModel }>(),
);

// PATCH
export const actionTaskPatch = createAction(
    '[Task] patch',
    props<{ task_id: string; data: TaskPayloadModel }>(),
);

// CHANGE ORDER
export const actionTaskChangeOrder = createAction(
    '[Task] change order',
    props<KanbanBoardChangeOrderPayload>(),
);

// Unified success action for create + patch
export const actionTaskUpsert = createAction('[Task] upsert', props<{ task: TaskModel }>());

// DELETE
export const actionTaskDelete = createAction('[Task] delete', props<{ task_id: string }>());
export const actionTaskDeleteSuccess = createAction(
    '[Task] delete success',
    props<{ taskId: string }>(),
);

// ERROR
export const actionTaskFailed = createAction('[Task] failed', props<{ error: any }>());
