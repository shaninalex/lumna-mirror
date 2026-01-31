import { ProjectModel, ProjectState, selectProjectByID } from '@entities/project';
import { ActivatedRouteSnapshot, ResolveFn, RouterStateSnapshot } from '@angular/router';
import { inject } from '@angular/core';
import { take, tap } from 'rxjs';
import { Store } from '@ngrx/store';

export const projectResolver: ResolveFn<ProjectModel | undefined> = (
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot,
) => {
    const store = inject(Store<ProjectState>);
    return store.select(selectProjectByID(route.params['id'])).pipe(
        take(1),
        tap((project) => {
            if (project) {
                // store.dispatch(StatusListGetActions({ projectId: project.id }));
                // store.dispatch(TaskListGetActions({ projectId: project.id }));
            }
        }),
    );
};
