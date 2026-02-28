import { ActivatedRouteSnapshot, ResolveFn } from '@angular/router';
import { Store } from '@ngrx/store';
import { inject } from '@angular/core';
import {filter, take, tap} from 'rxjs';
import {
    actionProjectList,
    ProjectModel,
    ProjectState,
    selectProjectByID,
} from '@entities/project';

export const projectResolver: ResolveFn<ProjectModel | null> = (route: ActivatedRouteSnapshot) => {
    const store = inject(Store<ProjectState>);
    const id = Number(route.paramMap.get('id'));

    store.dispatch(actionProjectList());

    return store.select(selectProjectByID(id)).pipe(
        filter(Boolean),
        take(1)
    );

    // const store = inject(Store<ProjectState>);
    // const projectId = Number(route.paramMap.get('id'));
    // return store.select(selectProjectByID(projectId)).pipe(
    //     tap((project) => {
    //         if (!project) {
    //             store.dispatch(actionProjectList());
    //         }
    //     }),
    //     filter((project) => !!project),
    // );
};
