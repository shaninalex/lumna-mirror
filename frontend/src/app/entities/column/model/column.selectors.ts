import { createFeatureSelector, createSelector } from '@ngrx/store';
import { columnAdapter, ColumnState } from './column.store';

const selectColumnFeature = createFeatureSelector<ColumnState>('column');

const columnSelectors = columnAdapter.getSelectors();

const selectColumns = createSelector(selectColumnFeature, (state) =>
    columnSelectors.selectAll(state),
);

// without sorting
export const selectColumnsByBoardIdNoSort = (boardId: number) =>
    createSelector(selectColumns, (columns) => columns.filter((b) => b.board_id === boardId));

// with sorting
export const selectColumnsByBoardId = (boardId: number) =>
    createSelector(selectColumns, (columns) =>
        columns.filter((b) => b.board_id === boardId).sort((a, b) => a.order - b.order),
    );

export const selectColumnsById = (columnId: number) =>
    createSelector(selectColumns, (columns) => columns.find((b) => b.id === columnId));
