import { Injectable, inject } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { catchError, exhaustMap, map, of } from 'rxjs';
import { actionKanban } from './kanban.actions';
import { KanbanApi } from '../api';
import { actionsColumns } from '@entities/column';
import type { HttpErrorResponse } from '@angular/common/http';
import { fromErrorResponse } from '@shared/models';

@Injectable()
export class KanbanEffects {
    private actions$ = inject(Actions);
    private kanbanApi = inject(KanbanApi);

    onDropColumnEffect$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionKanban.dropColumn),
            exhaustMap((action) =>
                this.kanbanApi.SetColumns(action.event).pipe(
                    map((columns) => actionsColumns.loadByBoardIdSuccess({ columns })),
                    catchError((err: HttpErrorResponse) =>
                        of(actionsColumns.reorderFailed({ errors: fromErrorResponse(err) })),
                    ),
                ),
            ),
        ),
    );
}
