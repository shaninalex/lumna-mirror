import { ColumnModel } from "@entities/column";
import { createAction, props } from "@ngrx/store";

export const actionKanbanLoadColumns = createAction(
    "[Kanban] load columns",
    props<{ boardId: number }>()
);

export const actionKanbanColumnsLoaded = createAction(
    "[Kanban] columns loaded",
    props<{ boardId: number; columns: ColumnModel[] }>()
);
