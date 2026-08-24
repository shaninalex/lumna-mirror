import { inject, Injectable } from "@angular/core";
import { Actions, createEffect, ofType } from "@ngrx/effects";
import { catchError, exhaustMap, map, of, switchMap } from "rxjs";
import { WorkspaceApi } from "../api/workspace.api";
import type { HttpErrorResponse } from "@angular/common/http";
import { fromErrorResponse } from "@shared/models";
import { Router } from "@angular/router";
import { actionWorkspace } from "./workspace.actions";

@Injectable()
export class WorkspaceEffects {
    private actions$ = inject(Actions);
    private api = inject(WorkspaceApi);
    private router = inject(Router);

    workspace_list$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionWorkspace.getList),
            exhaustMap(() => 
                this.api.List().pipe(
                    switchMap((list) => of(actionWorkspace.setList({ list })))
                )
            )
        )
    );

    workspace_create$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionWorkspace.create),
            exhaustMap((action) =>
                this.api.Create(action.data).pipe(
                    switchMap((workspace) => of(actionWorkspace.created({ data: workspace }))),
                    catchError((err: HttpErrorResponse) => of(actionWorkspace.createFailed({ errors: fromErrorResponse(err) })))
                )
            )
        )
    );

    workspace_created$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionWorkspace.created),
            map((action) => action.data.id),
            map((id) => this.router.navigate(["/app/w", id]))
        ),
        { dispatch: false }
    );

    // workspace_set_current$ = createEffect(() =>
    //     this.actions$.pipe(
    //         ofType(actionWorkspace.SetCurrent),
    //         map((action) => action.id),
    //         filter(id => id !== null),
    //         tap((id) => localStorage.setItem('last_workspace_id', id.toString())),
    //     ),
    //     { dispatch: false }
    // );
}
