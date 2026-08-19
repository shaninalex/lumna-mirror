import { createAction, props } from '@ngrx/store';
import type { BoardModel, BoardPayloadModel } from './board.model';
import type { Error } from '@shared/models';

// Single
export const actionBoardGet = createAction('[Board] get', props<{ boardId: number }>());
export const actionBoardSet = createAction('[Board] set', props<{ board: BoardModel }>());
export const actionBoardGetFailed = createAction('[Board] get failed', props<{ errors: Error[] }>());

// LIST
export const actionBoardGetList = createAction('[Board] get list', props<{ projectId: number }>());
export const actionBoardSetList = createAction(
    '[Board] set list',
    props<{ boards: BoardModel[] }>(),
);

// CREATE
export const actionBoardCreate = createAction(
    '[Board] create',
    props<{ data: BoardPayloadModel }>(),
);
export const actionBoardCreateSuccess = createAction(
    '[Board] create success',
    props<{ board: BoardModel }>(),
);
export const actionBoardCreateFailed = createAction(
    '[Board] create failed',
    props<{ errors: Error[] }>(),
);

// PATCH
export const actionBoardPatch = createAction(
    '[Board] patch',
    props<{ boardId: number; data: BoardPayloadModel }>(),
);
export const actionBoardPatchSuccess = createAction(
    '[Board] patch success',
    props<{ board: BoardModel }>(),
);
export const actionBoardPatchFailed = createAction(
    '[Board] patch failed',
    props<{ errors: Error[] }>(),
);

// DELETE
export const actionBoardDelete = createAction('[Board] delete', props<{ boardId: number }>());
export const actionBoardDeleteSuccess = createAction(
    '[Board] delete success',
    props<{ boardId: number }>(),
);
export const actionBoardDeleteFailed = createAction(
    '[Board] delete failed',
    props<{ errors: Error[] }>(),
);
