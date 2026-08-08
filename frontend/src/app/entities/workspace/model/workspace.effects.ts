import { inject, Injectable } from "@angular/core";
import { Actions, createEffect, ofType } from "@ngrx/effects";
import { catchError, exhaustMap, map, of, switchMap } from "rxjs";
import {
    actionWorkspaceCreate,
    actionWorkspaceCreateFailed,
    actionWorkspaceCreateSuccess,
    actionWorkspaceGetList,
    actionWorkspaceSetList
} from "./workspace.actions";
import { WorkspaceApi } from "../api/workspace.api";
import { HttpErrorResponse } from "@angular/common/http";
import { fromErrorResponse } from "@shared/models";

@Injectable()
export class WorkspaceEffects {
    private actions$ = inject(Actions);
    private api = inject(WorkspaceApi);

    workspace_list$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionWorkspaceGetList),
            exhaustMap(() =>
                this.api
                    .List()
                    .pipe(
                        switchMap((list) =>
                            of(actionWorkspaceSetList({ list }))
                        )
                    )
            )
        )
    );

    workspace_create$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionWorkspaceCreate),
            exhaustMap((action) =>
                this.api.Create(action.data).pipe(
                    switchMap((workspace) =>
                        of(
                            actionWorkspaceCreateSuccess({
                                data: workspace
                            })
                        )
                    ),
                    catchError((err: HttpErrorResponse) =>
                        of(
                            actionWorkspaceCreateFailed({
                                errors: fromErrorResponse(err)
                            })
                        )
                    )
                )
            )
        )
    );
}
