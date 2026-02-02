import { BoardModel } from '@entities/board';
import { createEntityAdapter, EntityState } from '@ngrx/entity';
import { createReducer, on } from '@ngrx/store';
import { actionBoardSetList, actionBoardUpsert, actionBoardDeleteSuccess } from './board.actions';

export interface BoardState extends EntityState<BoardModel> {}
export const boardAdapter = createEntityAdapter<BoardModel>();
const initialState = boardAdapter.getInitialState();

export const boardReducer = createReducer(
    initialState,
    on(actionBoardSetList, (state, { boards }) => boardAdapter.addMany(boards, state)),
    on(actionBoardUpsert, (state, { board }) => boardAdapter.upsertOne(board, state)),
    on(actionBoardDeleteSuccess, (state, { boardId }) => boardAdapter.removeOne(boardId, state)),
);
