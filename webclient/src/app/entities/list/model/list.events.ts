import {eventGroup} from '@ngrx/signals/events';
import {type} from '@ngrx/signals';
import {ListModel, ListPayloadModel} from '@entities/list';

export const listEvents = eventGroup({
    source: 'List',
    events: {
        getLists: type<number>(),
        create: type<{boardId: number, data: ListPayloadModel}>(),
        patch: type<{boardId: number, data: ListPayloadModel}>(),
        delete: type<number>(),

        setLists: type<ListModel[]>(),
        setList: type<ListModel>(),

        failed: type<any>(),

        // internal, do not used in components
        _patchSuccess: type<ListModel>(),
        _deleteSuccess: type<number>(),
    },
});
