import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';
import { Store } from '@ngrx/store';
import { map, switchMap, take } from 'rxjs/operators';
import { of } from 'rxjs';
import {
    actionProjectSetCurrent,
    actionProjectsSetList,
    ProjectApi,
    ProjectModel,
    selectProjectsByWorkspaceID,
} from '@entities/project';
import { selectCurrentWorkspaceId } from '@entities/workspace';
import { parseRouteId } from '@shared/utils';

function getOrFetchProjects(store: Store, projectApi: ProjectApi, workspaceId: number) {
    return store.select(selectProjectsByWorkspaceID(workspaceId)).pipe(
        take(1),
        switchMap((projects) => {
            if (projects && projects.length > 0) {
                return of(projects);
            }
            return projectApi.GetProjects(workspaceId).pipe(
                map((list: ProjectModel[]) => {
                    store.dispatch(actionProjectsSetList({ projects: list }));
                    return list;
                })
            );
        })
    );
}

export const activeProjectGuard: CanActivateFn = (route) => {
    const store = inject(Store);
    const router = inject(Router);
    const projectApi = inject(ProjectApi);
    const projectId = parseRouteId(route.paramMap.get('projectId'));

    // activeWorkspaceGuard runs on the parent route, so the current workspace is already set
    return store.select(selectCurrentWorkspaceId).pipe(
        take(1),
        switchMap((workspaceId) => {
            if (workspaceId === null) {
                return of(router.createUrlTree(['/app/workspaces']));
            }

            return getOrFetchProjects(store, projectApi, workspaceId).pipe(
                map((projects: ProjectModel[]) => {
                    if (!projects || projects.length === 0) {
                        // Nothing to open in this workspace, let the entry page take over
                        return router.createUrlTree(['/app/w', workspaceId]);
                    }

                    const targetProject =
                        projectId === null
                            ? undefined
                            : projects.find((p: ProjectModel) => p.id === projectId);

                    if (!targetProject) {
                        // Malformed or unknown project ID in URL
                        return router.createUrlTree(['/app/w', workspaceId]);
                    }

                    store.dispatch(actionProjectSetCurrent({ id: targetProject.id }));
                    return true;
                })
            );
        })
    );
};
