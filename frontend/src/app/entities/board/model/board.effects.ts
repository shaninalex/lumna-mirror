import { inject, Injectable } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';

import { BoardApi } from '../api/board.api';
import {
    actionBoardCreate,
    actionBoardDelete,
    actionBoardDeleteSuccess,
    actionBoardCreateFailed,
    actionBoardGet,
    actionBoardGetList,
    actionBoardPatch,
    actionBoardSetList,
    actionBoardCreateSuccess,
    actionBoardPatchSuccess,
    actionBoardPatchFailed,
    actionBoardDeleteFailed,
    actionBoardSet,
    actionBoardGetFailed,
} from './board.actions';
import { catchError, exhaustMap, of, switchMap } from 'rxjs';
import type { HttpErrorResponse } from '@angular/common/http';
import { fromErrorResponse } from '@shared/models';

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
                this.boardApi.Create(action.data).pipe(
                    switchMap((board) => of(actionBoardCreateSuccess({ board }))),
                    catchError((err: HttpErrorResponse) =>
                        of(actionBoardCreateFailed({ errors: fromErrorResponse(err) })),
                    ),
                ),
            ),
        ),
    );

    patch_board$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionBoardPatch),
            exhaustMap((action) =>
                this.boardApi.Patch(action.boardId, action.data).pipe(
                    switchMap((board) => of(actionBoardPatchSuccess({ board }))),
                    catchError((err: HttpErrorResponse) =>
                        of(actionBoardPatchFailed({ errors: fromErrorResponse(err) })),
                    ),
                ),
            ),
        ),
    );

    delete_board$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionBoardDelete),
            exhaustMap((action) =>
                this.boardApi.Delete(action.boardId).pipe(
                    switchMap(() => of(actionBoardDeleteSuccess({ boardId: action.boardId }))),
                    catchError((err: HttpErrorResponse) =>
                        of(actionBoardDeleteFailed({ errors: fromErrorResponse(err) })),
                    ),
                ),
            ),
        ),
    );

    get_board$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionBoardGet),
            exhaustMap((action) =>
                this.boardApi.Get(action.boardId).pipe(
                    switchMap((board) => of(actionBoardSet({ board }))),
                    catchError((err: HttpErrorResponse) =>
                        of(actionBoardGetFailed({ errors: fromErrorResponse(err) })),
                    ),
                ),
            ),
        ),
    );
}
