import { inject, Injectable } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';

import { ListApi } from '../api/list.api';

import { catchError, exhaustMap, of, switchMap } from 'rxjs';
import type { HttpErrorResponse } from '@angular/common/http';
import { fromErrorResponse } from '@shared/models';
import {
    actionListCreate,
    actionListCreateFailed,
    actionListCreateSuccess,
    actionListDelete,
    actionListDeleteFailed,
    actionListDeleteSuccess,
    actionListGet,
    actionListGetFailed,
    actionListGetList,
    actionListPatch,
    actionListPatchFailed,
    actionListPatchSuccess,
    actionListSet,
    actionListSetList,
} from './list.actions';

@Injectable()
export class ListsEffects {
    private actions$ = inject(Actions);
    private listApi = inject(ListApi);

    list_lists$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionListGetList),
            exhaustMap((action) =>
                this.listApi
                    .List(action.projectId)
                    .pipe(switchMap((lists) => of(actionListSetList({ lists: lists })))),
            ),
        ),
    );

    create_list$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionListCreate),
            exhaustMap((action) =>
                this.listApi.Create(action.data).pipe(
                    switchMap((list) => of(actionListCreateSuccess({ list }))),
                    catchError((err: HttpErrorResponse) =>
                        of(actionListCreateFailed({ errors: fromErrorResponse(err) })),
                    ),
                ),
            ),
        ),
    );

    patch_list$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionListPatch),
            exhaustMap((action) =>
                this.listApi.Patch(action.listId, action.data).pipe(
                    switchMap((list) => of(actionListPatchSuccess({ list }))),
                    catchError((err: HttpErrorResponse) =>
                        of(actionListPatchFailed({ errors: fromErrorResponse(err) })),
                    ),
                ),
            ),
        ),
    );

    delete_list$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionListDelete),
            exhaustMap((action) =>
                this.listApi.Delete(action.listId).pipe(
                    switchMap(() => of(actionListDeleteSuccess({ listId: action.listId }))),
                    catchError((err: HttpErrorResponse) =>
                        of(actionListDeleteFailed({ errors: fromErrorResponse(err) })),
                    ),
                ),
            ),
        ),
    );

    get_list$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionListGet),
            exhaustMap((action) =>
                this.listApi.Get(action.listId).pipe(
                    switchMap((list) => of(actionListSet({ list }))),
                    catchError((err: HttpErrorResponse) =>
                        of(actionListGetFailed({ errors: fromErrorResponse(err) })),
                    ),
                ),
            ),
        ),
    );
}
