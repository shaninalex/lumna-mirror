import {eventGroup} from '@ngrx/signals/events';
import {type} from '@ngrx/signals';
import {BoardModel, BoardPayloadModel} from '@entities/board';

export const boardEvents = eventGroup({
    source: 'Board',
    events: {
        getList: type<number>(),
        create: type<BoardPayloadModel>(),
        patch: type<{projectId: number, data: BoardPayloadModel}>(),
        delete: type<number>(),

        setList: type<BoardModel[]>(),
        set: type<BoardModel>(),

        failed: type<any>(), // TODO: proper error typing

        // internal, do not used in components
        _patchSuccess: type<BoardModel>(),
        _deleteSuccess: type<number>(),
    },
});
