import { inject, Injectable } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';

import { BoardApi } from '../api/board.api';
import {
    actionBoardCreate,
    actionBoardDelete,
    actionBoardDeleteSuccess,
    actionBoardGetList,
    actionBoardPatch,
    actionBoardSetList,
    actionBoardUpsert,
} from './board.actions';
import { exhaustMap, of, switchMap } from 'rxjs';

@Injectable()
export class BoardsEffects {
    private actions$ = inject(Actions);
    private boardApi = inject(BoardApi);

    list_boards$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionBoardGetList),
            exhaustMap((action) =>
                this.boardApi
                    .List(action.projectId)
                    .pipe(switchMap((boards) => of(actionBoardSetList({ boards: boards })))),
            ),
        ),
    );

    create_board$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionBoardCreate),
            exhaustMap((action) =>
                this.boardApi
                    .Create(action.data)
                    .pipe(switchMap((board) => of(actionBoardUpsert({ board })))),
            ),
        ),
    );

    patch_board$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionBoardPatch),
            exhaustMap((action) =>
                this.boardApi
                    .Patch(action.boardId, action.data)
                    .pipe(switchMap((board) => of(actionBoardUpsert({ board })))),
            ),
        ),
    );

    delete_board$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionBoardDelete),
            exhaustMap((action) =>
                this.boardApi
                    .Delete(action.boardId)
                    .pipe(
                        switchMap(() => of(actionBoardDeleteSuccess({ boardId: action.boardId }))),
                    ),
            ),
        ),
    );
}
