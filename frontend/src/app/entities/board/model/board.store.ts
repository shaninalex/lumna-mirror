import type { BoardModel } from '@entities/board';
import type { EntityState } from '@ngrx/entity';
import { createEntityAdapter } from '@ngrx/entity';
import { createReducer, on } from '@ngrx/store';
import {
    actionBoardSetList,
    actionBoardDeleteSuccess,
    actionBoardSet,
    actionBoardCreateSuccess,
    actionBoardPatchSuccess,
} from './board.actions';

export type BoardState = EntityState<BoardModel>;
export const boardAdapter = createEntityAdapter<BoardModel>();
const initialState = boardAdapter.getInitialState();

export const boardReducer = createReducer(
    initialState,
    on(actionBoardSetList, (state, { boards }) => boardAdapter.addMany(boards, state)),
    on(actionBoardSet, (state, { board }) => boardAdapter.upsertOne(board, state)),
    on(actionBoardCreateSuccess, (state, { board }) => boardAdapter.upsertOne(board, state)),
    on(actionBoardPatchSuccess, (state, { board }) => boardAdapter.upsertOne(board, state)),
    on(actionBoardDeleteSuccess, (state, { boardId }) => boardAdapter.removeOne(boardId, state)),
);
