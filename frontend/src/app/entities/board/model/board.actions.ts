import { createAction, props } from '@ngrx/store';
import { BoardModel, BoardPayloadModel } from '@entities/board';

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

// PATCH
export const actionBoardPatch = createAction(
    '[Board] patch',
    props<{ boardId: number; data: BoardPayloadModel }>(),
);

// SET single board:
export const actionBoardUpsert = createAction('[Board] upsert', props<{ board: BoardModel }>());

// DELETE
export const actionBoardDelete = createAction('[Board] delete', props<{ boardId: number }>());

// to remove from store
export const actionBoardDeleteSuccess = createAction(
    '[Board] delete success',
    props<{ boardId: number }>(),
);

// ERROR
export const actionBoardFailed = createAction('[Board] failed', props<{ error: any }>());
