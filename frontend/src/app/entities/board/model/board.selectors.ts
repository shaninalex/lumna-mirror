import { createFeatureSelector, createSelector } from '@ngrx/store';
import { boardAdapter, BoardState } from './board.store';

export const selectBoardFeature = createFeatureSelector<BoardState>('board');

export const boardSelectors = boardAdapter.getSelectors();

export const selectBoards = createSelector(selectBoardFeature, (state) =>
    boardSelectors.selectAll(state),
);

/**
 * Select boards by project id, always sorted by `order` ASC
 */
export const selectBoardsByProjectId = (projectId: string) =>
    createSelector(selectBoards, (boards) => boards.filter((b) => b.project_id === projectId));

export const selectBoardById = (boardId: string) =>
    createSelector(
        selectBoardFeature,
        (state) => boardSelectors.selectEntities(state)[boardId] ?? null,
    );
