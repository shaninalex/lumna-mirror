import { ColumnModel } from './column.model';
import { createEntityAdapter, EntityState, Update } from '@ngrx/entity';
import { createReducer, on } from '@ngrx/store';
import {
    actionColumnSetList,
    actionColumnUpsert,
    actionColumnDeleteSuccess,
    actionColumnChangeOrder,
} from './column.actions';

export interface ColumnState extends EntityState<ColumnModel> {}
export const columnAdapter = createEntityAdapter<ColumnModel>();
const initialState = columnAdapter.getInitialState();

export const columnReducer = createReducer(
    initialState,

    // Replace all columns for a board (after load)
    on(actionColumnSetList, (state, { columns }) => columnAdapter.setAll(columns, state)),

    // Create OR update a column
    on(actionColumnUpsert, (state, { column }) => columnAdapter.upsertOne(column, state)),

    // Remove column after delete success
    on(actionColumnDeleteSuccess, (state, { columnId }) =>
        columnAdapter.removeOne(columnId, state),
    ),

    // Change order for multiple columns
    on(actionColumnChangeOrder, (state, { columns }) =>
        columnAdapter.updateMany(
            columns.map(
                ({ id, order }): Update<ColumnModel> => ({
                    id,
                    changes: { order },
                }),
            ),
            state,
        ),
    ),
);
