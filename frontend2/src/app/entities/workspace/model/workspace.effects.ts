import { inject, Injectable } from "@angular/core";
import { Actions, createEffect, ofType } from "@ngrx/effects";
import { exhaustMap, map, of, switchMap } from "rxjs";
import {
    actionWorkspaceGetList,
    actionWorkspaceSetList
} from "./workspace.actions";
import { WorkspaceModel } from "./workspace.model";

@Injectable()
export class WorkspaceEffects {
    private actions$ = inject(Actions);

    workspace_list$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionWorkspaceGetList),
            exhaustMap(() => {
                const list: WorkspaceModel[] = [
                    {
                        id: 1,
                        slug: "lumna-1",
                        title: "Lumna",
                        icon: "/img/project.svg"
                    }
                ];

                return of(actionWorkspaceSetList({ list }));
            })
        )
    );
}
