import {ResolveFn} from '@angular/router';
import {inject} from '@angular/core';
import {Store} from '@ngrx/store';
import {AppState} from '@client/shared/store';
import {selectProject} from '@client/entities/project/model/project.selectors';
import {Project} from '@client/entities/project';
import {take, tap} from 'rxjs';
import {GetStatusListActions} from '@client/entities/status';
import {GetTasksActions} from '@client/entities/task';

export const projectResolver: ResolveFn<Project | undefined> = (route) => {
    const store: Store<AppState> = inject(Store<AppState>)
    return store.select(selectProject(route.params["projectKey"])).pipe(
        take(1),
        tap(project => {
            if (project) {
                store.dispatch(GetStatusListActions({projectId: project.id}))
                store.dispatch(GetTasksActions({projectId: project.id}))
            }
        })
    );
};
