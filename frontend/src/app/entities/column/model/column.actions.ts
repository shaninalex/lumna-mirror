import { createActionGroup, props } from '@ngrx/store';
import type { ColumnPayloadModel, ColumnModel } from './column.model';
import type { Error } from '@shared/models';

export const actionsColumns = createActionGroup({
    source: 'Column',
    events: {
        'load by board id': props<{ board_id: number }>(),
        'load by board id success': props<{ columns: ColumnModel[] }>(),
        'load by board id failed': props<{ errors: Error[] }>(),
        create: props<{ payload: ColumnPayloadModel }>(),
        'create success': props<{ column: ColumnModel }>(),
        'create failed': props<{ errors: Error[] }>(),
        'reorder failed': props<{ errors: Error[] }>(),
    },
});
