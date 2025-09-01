import {ResolveFn, Routes} from '@angular/router';
import {PrimaryRoot} from './primary-root';
import {Overview} from './overview/overview';
import {overviewResolver} from './overview/overview.resolver';
import {ProjectsList} from './projects-list/projects-list'
import {projectListResolver} from './projects-list/projects-list.resolver';
import {GetSessionAction} from '@client/entities/session';
import {Store} from '@ngrx/store';
import {inject} from '@angular/core';
import {AppState} from '@client/shared/store';
import {ProjectService} from '@client/entities/project/api/project.service';
import {TaskService} from '@client/entities/task/api/task.service';

export const sessionResolver: ResolveFn<void> = () => {
    const store = inject(Store<AppState>);
    store.dispatch(GetSessionAction());
    return undefined;
};

export const mainRoutes: Routes = [
    {
        path: "",
        component: PrimaryRoot,
        resolve: { session: sessionResolver },
        // providers: [ProjectService, TaskService],
        children: [
            {
                path: "",
                component: Overview,
                resolve: { overview: overviewResolver }
            },
            {
                path: "projects",
                component: ProjectsList,
                resolve: { projects: projectListResolver }
            }
        ]
    }
];
