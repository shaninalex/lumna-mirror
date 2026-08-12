import { Actions, createEffect, ofType } from '@ngrx/effects';
import { inject, Injectable } from '@angular/core';
import { filter, map } from 'rxjs';
import { actionSessionAuthenticated } from './session.actions';
import { actionWorkspaceGetList, actionWorkspaceSetCurrent } from '@entities/workspace';
import { actionProjectList } from '@entities/project';

@Injectable()
export class AppEffects {
    private actions$ = inject(Actions);

    authenticated_successfull$ = createEffect(() => 
        this.actions$.pipe(
            ofType(actionSessionAuthenticated),
            map(() => actionWorkspaceGetList()),
        )
    );

    setCurrentWorkspace$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionWorkspaceSetCurrent),
            map((action) => action.id),
            filter((id) => id !== null),
            map((id) => actionProjectList({ workspace_id: id}))
        )
    )
}
