import { inject, Injectable } from "@angular/core";
import { Actions, createEffect, ofType } from "@ngrx/effects";
import { EMPTY, exhaustMap, of, switchMap } from "rxjs";
import {
    actionWorkspaceGetList,
    actionWorkspaceSetList
} from "./workspace.actions";
import { routerNavigatedAction } from "@ngrx/router-store";
import { WorkspaceApi } from "../api";

@Injectable()
export class WorkspaceEffects {
    private actions$ = inject(Actions);
    private api = inject(WorkspaceApi);

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
}
