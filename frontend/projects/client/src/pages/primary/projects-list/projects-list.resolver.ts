import {ResolveFn} from '@angular/router';
import {inject} from '@angular/core';
import {Store} from '@ngrx/store';
import {AppState} from '@client/shared/store';
import {selectProjects} from '@client/entities/project/model/project.selectors';
import {Project, SetProjectsAction} from '@client/entities/project';
import {of, switchMap, take, tap} from 'rxjs';
import {ProjectService} from '@client/entities/project/api/project.service';

export const projectListResolver: ResolveFn<Array<Project> | undefined> = (route) => {
    const store: Store<AppState> = inject(Store<AppState>)
    const api = inject(ProjectService)

    return store.select(selectProjects).pipe(
        take(1),
        switchMap(projects => {
            if (!projects.length) {
                return api.List().pipe(
                    tap(projects => store.dispatch(SetProjectsAction({payload: projects})))
                );
            }
            return of(projects);
        })
    );
};
