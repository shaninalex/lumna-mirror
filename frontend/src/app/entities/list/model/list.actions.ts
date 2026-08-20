import { createAction, props } from '@ngrx/store';
import type { ListModel, ListPayloadModel } from './list.model';
import type { Error } from '@shared/models';

// Single
export const actionListGet = createAction('[List] get', props<{ listId: number }>());
export const actionListSet = createAction('[List] set', props<{ list: ListModel }>());
export const actionListGetFailed = createAction('[List] get failed', props<{ errors: Error[] }>());

// LIST
export const actionListGetList = createAction('[List] get list', props<{ projectId: number }>());
export const actionListSetList = createAction(
    '[List] set list',
    props<{ lists: ListModel[] }>(),
);

// CREATE
export const actionListCreate = createAction(
    '[List] create',
    props<{ data: ListPayloadModel }>(),
);
export const actionListCreateSuccess = createAction(
    '[List] create success',
    props<{ list: ListModel }>(),
);
export const actionListCreateFailed = createAction(
    '[List] create failed',
    props<{ errors: Error[] }>(),
);

// PATCH
export const actionListPatch = createAction(
    '[List] patch',
    props<{ listId: number; data: ListPayloadModel }>(),
);
export const actionListPatchSuccess = createAction(
    '[List] patch success',
    props<{ list: ListModel }>(),
);
export const actionListPatchFailed = createAction(
    '[List] patch failed',
    props<{ errors: Error[] }>(),
);

// DELETE
export const actionListDelete = createAction('[List] delete', props<{ listId: number }>());
export const actionListDeleteSuccess = createAction(
    '[List] delete success',
    props<{ listId: number }>(),
);
export const actionListDeleteFailed = createAction(
    '[List] delete failed',
    props<{ errors: Error[] }>(),
);
