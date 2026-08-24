import { createActionGroup, props } from '@ngrx/store';
import type { BoardModel, BoardPayloadModel } from './board.model';
import type { Error } from '@shared/models';

export const actionBoard = createActionGroup({
    source: 'Board',
    events: {
        'get by id': props<{ boardId: number }>(),
        set: props<{ board: BoardModel }>(),
        'get failed': props<{ errors: Error[] }>(),

        'get list': props<{ projectId: number }>(),
        'set list': props<{ boards: BoardModel[] }>(),

        create: props<{ data: BoardPayloadModel }>(),
        'create success': props<{ board: BoardModel }>(),
        'create failed': props<{ errors: Error[] }>(),

        patch: props<{ boardId: number; data: BoardPayloadModel }>(),
        'patch success': props<{ board: BoardModel }>(),
        'patch failed': props<{ errors: Error[] }>(),

        delete: props<{ boardId: number }>(),
        'delete success': props<{ boardId: number }>(),
        'delete failed': props<{ errors: Error[] }>(),
    },
});
