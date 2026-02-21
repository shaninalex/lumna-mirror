import { ActivatedRouteSnapshot, ResolveFn, RouterStateSnapshot } from '@angular/router';
import { Store } from '@ngrx/store';
import { inject } from '@angular/core';
import { filter, tap } from 'rxjs';
import {
    actionProjectList,
    ProjectModel,
    ProjectState,
    selectProjectByID,
} from '@entities/project';
import { UiService } from '@shared/ui';

export const projectResolver: ResolveFn<ProjectModel | null> = (
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot,
) => {
    const store = inject(Store<ProjectState>);
    const ui = inject(UiService);
    const projectId = Number(route.paramMap.get('id'));

    return store.select(selectProjectByID(projectId)).pipe(
        tap((project) => {
            if (!project) {
                store.dispatch(actionProjectList());
            }
        }),
        filter((project) => !!project),
        tap((project) => ui.setPageTitle(`Project: ${project.title}`)),
    );
};
