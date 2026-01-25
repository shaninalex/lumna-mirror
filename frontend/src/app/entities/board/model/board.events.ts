import { eventGroup } from '@ngrx/signals/events';
import { type } from '@ngrx/signals';
import { BoardModel, BoardPayloadModel } from '@entities/board';

export const boardEvents = eventGroup({
    source: 'Board',
    events: {
        getList: type<string>(),
        create: type<BoardPayloadModel>(),
        patch: type<{ boardId: string; data: BoardPayloadModel }>(),
        delete: type<string>(),

        setList: type<BoardModel[]>(),
        set: type<BoardModel>(),

        failed: type<any>(), // TODO: proper error typing

        // internal, do not used in components
        _patchSuccess: type<BoardModel>(),
        _deleteSuccess: type<string>(),
    },
});
