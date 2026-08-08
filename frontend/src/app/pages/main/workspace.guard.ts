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
    WorkspaceModel
} from '@entities/workspace';

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
                })
            );
        })
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
            const targetWorkspace =
                workspaces.find((w: WorkspaceModel) => w.id.toString() === savedId) || workspaces[0];

            return router.createUrlTree(['/app/w', targetWorkspace.id, 'inbox']);
        })
    );
};

export const activeWorkspaceGuard: CanActivateFn = (route) => {
    const store = inject(Store);
    const router = inject(Router);
    const workspaceIdParam = route.paramMap.get('workspaceId');

    return getOrFetchWorkspaces().pipe(
        map((workspaces: WorkspaceModel[]) => {
            if (!workspaces || workspaces.length === 0) {
                return router.createUrlTree(['/app/workspaces/create']);
            }

            const numericId = workspaceIdParam ? parseInt(workspaceIdParam, 10) : NaN;
            const targetWorkspace = workspaces.find((w: WorkspaceModel) => w.id === numericId);

            if (!targetWorkspace) {
                // If invalid workspace ID in URL, fallback to first available workspace
                return router.createUrlTree(['/app/w', workspaces[0].id, 'inbox']);
            }

            store.dispatch(actionWorkspaceSetCurrent({ id: targetWorkspace.id }));
            return true;
        })
    );
};

