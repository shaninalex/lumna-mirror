import { createFeatureSelector, createSelector } from '@ngrx/store';
import type { BoardState } from './board.store';
import { boardAdapter } from './board.store';

export const selectBoardFeature = createFeatureSelector<BoardState>('board');

export const boardSelectors = boardAdapter.getSelectors();

export const selectBoards = createSelector(selectBoardFeature, (state) =>
    boardSelectors.selectAll(state),
);

/**
 * Select boards by project id, always sorted by `order` ASC
 */
export const selectBoardsByProjectId = (projectId: number) =>
    createSelector(selectBoards, (boards) => {
        return boards.filter((b) => b.project_id === projectId);
    });

export const selectBoardById = (boardId: number) =>
    createSelector(
        selectBoardFeature,
        (state) => boardSelectors.selectEntities(state)[boardId] ?? null,
    );
