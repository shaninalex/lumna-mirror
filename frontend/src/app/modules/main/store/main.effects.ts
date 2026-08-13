import { Actions, createEffect, ofType } from '@ngrx/effects';
import { inject, Injectable } from '@angular/core';
import { filter, map, tap } from 'rxjs';
import { ROUTER_NAVIGATED } from '@ngrx/router-store';
import { actionSessionAuthenticated } from '@core';
import { actionProjectList, actionProjectSetCurrent } from '@entities/project';
import { actionWorkspaceGetList, actionWorkspaceSetCurrent } from '@entities/workspace';
import { LocalStorageService } from '@shared/services';
import { actionTaskGetList } from '@entities/task';

@Injectable()
export class MainEffects {
    private actions$ = inject(Actions);
    private localStorageService = inject(LocalStorageService);

    routeChanged$ = createEffect(
        () =>
            this.actions$.pipe(
                ofType(ROUTER_NAVIGATED),
                map((action) => action.payload.event.url),
                filter((url: string) => url.startsWith('/app')),
                tap((url) => this.localStorageService.set('last_url', url)),
            ),
        { dispatch: false },
    );

    authenticated_successfull$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionSessionAuthenticated),
            map(() => actionWorkspaceGetList()),
        ),
    );

    setCurrentWorkspace$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionWorkspaceSetCurrent),
            map((action) => action.id),
            filter((id) => id !== null),
            map((id) => actionProjectList({ workspace_id: id })),
        ),
    );

    loadProjectTasks$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionProjectSetCurrent),
            map((action) => action.id),
            filter((id) => id !== null),
            map((id) => actionTaskGetList({ query: { project_id: id } })),
        ),
    );
}
