import { createAction, props } from '@ngrx/store';
import { ColumnModel, ColumnPayloadModel } from './column.model';

// LIST
export const actionColumnGetList = createAction('[Column] get list', props<{ boardId: string }>());

export const actionColumnSetList = createAction(
    '[Column] set list',
    props<{ columns: ColumnModel[] }>(),
);

// CREATE
export const actionColumnCreate = createAction(
    '[Column] create',
    props<{ boardId: string; data: ColumnPayloadModel }>(),
);

// PATCH
export const actionColumnPatch = createAction(
    '[Column] patch',
    props<{ columnId: string; data: ColumnPayloadModel }>(),
);

// CHANGE ORDER
export const actionColumnChangeOrder = createAction(
    '[Column] change order',
    props<{ columns: Array<{ id: string; order: number }> }>(),
);

// DELETE
export const actionColumnDelete = createAction('[Column] delete', props<{ columnId: string }>());

// REPLACEMENT:
// - setList
// - _patchSuccess
// Unified success action for create + patch
export const actionColumnUpsert = createAction('[Column] upsert', props<{ column: ColumnModel }>());

// REPLACEMENT:
// - _deleteSuccess
export const actionColumnDeleteSuccess = createAction(
    '[Column] delete success',
    props<{ columnId: string }>(),
);

// ERROR
export const actionColumnFailed = createAction('[Column] failed', props<{ error: any }>());
