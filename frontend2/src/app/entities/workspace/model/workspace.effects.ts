import { inject, Injectable } from "@angular/core";
import { Actions, createEffect, ofType } from "@ngrx/effects";
import { catchError, EMPTY, exhaustMap, of, switchMap } from "rxjs";
import {
    actionWorkspaceCreate,
    actionWorkspaceCreateFailed,
    actionWorkspaceCreateSuccess,
    actionWorkspaceGetList,
    actionWorkspaceSetList
} from "./workspace.actions";
import { routerNavigatedAction } from "@ngrx/router-store";
import { WorkspaceApi } from "../api";
import { HttpErrorResponse } from "@angular/common/http";
import { fromErrorResponse } from "@shared/models";

@Injectable()
export class WorkspaceEffects {
    private api = inject(WorkspaceApi);
    private actions$ = inject(Actions);

    workspace_route_data$ = createEffect(
        () =>
            this.actions$.pipe(
                ofType(routerNavigatedAction),
                exhaustMap((action) => {
                    if (action.payload.routerState.url.includes("app")) {
                        console.log(action);
                        return EMPTY;
                    }
                    return EMPTY;
                })
            ),
        { dispatch: false }
    );

    workspaces_route_data$ = createEffect(() =>
        this.actions$.pipe(
            ofType(routerNavigatedAction),
            exhaustMap((action) => {
                if (
                    action.payload.routerState.url === "/workspaces" ||
                    action.payload.routerState.url.includes("app")
                ) {
                    return of(actionWorkspaceGetList());
                }
                return EMPTY;
            })
        )
    );

    workspace_list$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionWorkspaceGetList),
            exhaustMap(() =>
                this.api
                    .list()
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
                this.api.create(action.data).pipe(
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
