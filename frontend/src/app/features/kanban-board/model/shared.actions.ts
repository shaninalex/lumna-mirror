import { ColumnModel } from '@entities/column';
import { TaskModel } from '@entities/task';
import { createAction, props } from '@ngrx/store';

export const actionKanbanLoadColumns = createAction(
    '[Kanban] load columns',
    props<{ boardId: string }>(),
);

export const actionKanbanColumnsLoaded = createAction(
    '[Kanban] columns loaded',
    props<{ columns: ColumnModel[]; tasks: TaskModel[] }>(),
);
