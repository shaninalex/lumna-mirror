import { eventGroup } from '@ngrx/signals/events';
import { type } from '@ngrx/signals';
import { ColumnModel, ColumnPayloadModel } from './column.model';

export const listEvents = eventGroup({
    source: 'List',
    events: {
        getLists: type<string>(),
        create: type<{ boardId: string; data: ColumnPayloadModel }>(),
        patch: type<{ listId: string; data: ColumnPayloadModel }>(),
        changeOrder: type<{ lists: Array<{ id: string; order: number }> }>(),
        delete: type<string>(),

        setLists: type<ColumnModel[]>(),
        setList: type<ColumnModel>(),

        failed: type<any>(),

        // internal, do not used in components
        _patchSuccess: type<ColumnModel>(),
        _deleteSuccess: type<string>(),
    },
});
