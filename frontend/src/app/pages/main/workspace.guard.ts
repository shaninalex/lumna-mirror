import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';
import { Store } from '@ngrx/store';
import { map, switchMap, take } from 'rxjs/operators';
import { of } from 'rxjs';
import {
    actionWorkspaceSetCurrent,
    actionWorkspaceSetList,
    selectWorkspaceList,
    WorkspaceApi,
    WorkspaceModel,
} from '@entities/workspace';
import { parseRouteId } from '@shared/utils';

function getOrFetchWorkspaces() {
    const store = inject(Store);
    const workspaceApi = inject(WorkspaceApi);

    return store.select(selectWorkspaceList).pipe(
        take(1),
        switchMap((workspaces) => {
            if (workspaces && workspaces.length > 0) {
                return of(workspaces);
            }
            return workspaceApi.List().pipe(
                map((list: WorkspaceModel[]) => {
                    store.dispatch(actionWorkspaceSetList({ list }));
                    return list;
                }),
            );
        }),
    );
}

export const workspaceRedirectGuard: CanActivateFn = () => {
    const router = inject(Router);

    return getOrFetchWorkspaces().pipe(
        map((workspaces: WorkspaceModel[]) => {
            if (!workspaces || workspaces.length === 0) {
                return router.createUrlTree(['/app/workspaces/create']);
            }

            const savedId = localStorage.getItem('last_workspace_id');
            const targetWorkspace = workspaces.find(
                (w: WorkspaceModel) => w.id.toString() === savedId,
            ); // || workspaces[0];
            console.log('workspaceRedirectGuard: ', targetWorkspace);
            if (!targetWorkspace) {
                return router.createUrlTree(['/app/workspaces']);
            }

            return router.createUrlTree(['/app/w', targetWorkspace.id]);
        }),
    );
};

export const activeWorkspaceGuard: CanActivateFn = (route) => {
    const store = inject(Store);
    const router = inject(Router);
    const workspaceId = parseRouteId(route.paramMap.get('workspaceId'));

    return getOrFetchWorkspaces().pipe(
        map((workspaces: WorkspaceModel[]) => {
            if (!workspaces || workspaces.length === 0) {
                return router.createUrlTree(['/app/workspaces/create']);
            }

            const targetWorkspace =
                workspaceId === null
                    ? undefined
                    : workspaces.find((w: WorkspaceModel) => w.id === workspaceId);

            if (!targetWorkspace) {
                // Malformed or unknown workspace ID in URL
                return router.createUrlTree(['/app/workspaces']);
            }

            store.dispatch(actionWorkspaceSetCurrent({ id: targetWorkspace.id }));
            return true;
        }),
    );
};
