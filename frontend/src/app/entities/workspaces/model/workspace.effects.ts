import { inject, Injectable } from "@angular/core";
import { Actions, createEffect, ofType } from "@ngrx/effects";
import { exhaustMap, of, switchMap } from "rxjs";
import {
    actionWorkspaceGetList,
    actionWorkspaceSetList
} from "./workspace.actions";
import { WorkspaceApi } from "../api/workspace.api";

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
}
