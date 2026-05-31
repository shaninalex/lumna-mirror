import { inject, Injectable } from "@angular/core";
import { Actions, createEffect, ofType } from "@ngrx/effects";
import { catchError, exhaustMap, map, of, switchMap, withLatestFrom } from "rxjs";
import {
    actionWorkspaceCreate,
    actionWorkspaceCreateFailed,
    actionWorkspaceCreateSuccess,
    actionWorkspaceListRequested,
    actionWorkspaceSetCurrent,
    actionWorkspaceSetList
} from "./workspace.actions";
import { WorkspaceApi } from "../api";
import { HttpErrorResponse } from "@angular/common/http";
import { fromErrorResponse } from "@shared/models";
import { Store } from "@ngrx/store";
import { selectWorkspaceListLoaded } from "@entities/workspace/model/workspace.selectors";
import { filter } from "rxjs/operators";
import { routerNavigatedAction } from "@ngrx/router-store";

@Injectable()
export class WorkspaceEffects {
    private store = inject(Store);
    private api = inject(WorkspaceApi);
    private actions$ = inject(Actions);

    // TODO: get archived workspaces
    // /app/workspaces/archived

    routerEffect$ = createEffect(() =>
        this.actions$.pipe(
            ofType(routerNavigatedAction),
            map((action) =>
                findRouteParam(action.payload.routerState.root, "workspace-id")
            ),
            filter((id) => id !== null),
            map((id) => {
                return actionWorkspaceSetCurrent({ id: parseInt(id) });
            })
        )
    );

    workspace_list$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionWorkspaceListRequested),
            withLatestFrom(this.store.select(selectWorkspaceListLoaded)),
            filter(([_, loaded]) => !loaded),
            switchMap(() =>
                this.api
                    .list(true)
                    .pipe(map((list) => actionWorkspaceSetList({ list })))
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

function findRouteParam(route: any, paramName: string): string | null {
    let current = route;

    while (current) {
        if (current.params?.[paramName]) {
            return current.params[paramName];
        }

        current = current.firstChild;
    }

    return null;
}
