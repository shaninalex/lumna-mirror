import { Actions, createEffect, ofType } from '@ngrx/effects';
import { inject, Injectable } from '@angular/core';
import { filter, map, tap } from 'rxjs';
import { ROUTER_NAVIGATED, RouterNavigatedPayload } from '@ngrx/router-store';
import { actionSessionAuthenticated } from '@core';
import { actionProjectList } from '@entities/project';
import { actionWorkspaceGetList, actionWorkspaceSetCurrent } from '@entities/workspace';
import { LocalStorageService } from '@shared/services';

@Injectable()
export class MainEffects {
    private actions$ = inject(Actions);
    private localStorageService = inject(LocalStorageService);

    routeChanged$ = createEffect(
        () =>
            this.actions$.pipe(
                ofType(ROUTER_NAVIGATED),
                map((action) => action.payload),
                tap((payload: RouterNavigatedPayload) => this.localStorageService.set("last_url", payload.event.url)),
            ),
        { dispatch: false },
    );

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
