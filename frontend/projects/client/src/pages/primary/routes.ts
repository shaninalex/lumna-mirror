import {ResolveFn, Routes} from '@angular/router';
import {PrimaryRoot} from './primary-root';
import {Overview} from './overview/overview';
import {overviewResolver} from './overview/overview.resolver';
import {ProjectsListPageComponent} from './projects-list/projects-list-page.component'
import {projectListResolver} from './projects-list/projects-list.resolver';
import {Store} from '@ngrx/store';
import {inject} from '@angular/core';
import {AppState} from '@client/shared/store';
import {GetSessionAction} from '@client/entities/auth';
import {CanMatchPrimarySection} from '@client/pages/primary/guard';
import {AccountPageComponent} from './account-page/account-page';
import {ProjectDetailPageComponent} from '@client/pages/primary/project-detail';
import {ProjectsRootComponent} from '@client/pages/primary/projects-root';
import {SettingsPageComponent} from '@client/pages/primary/settings-page';

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
        canMatch: [CanMatchPrimarySection],
        children: [
            {
                path: "",
                component: Overview,
                resolve: { overview: overviewResolver }
            },
            {
                path: "projects",
                component: ProjectsRootComponent,
                resolve: { projects: projectListResolver },
                children: [
                    {
                        path: "",
                        component: ProjectsListPageComponent,
                    },
                    {
                        path: ":projectKey",
                        component: ProjectDetailPageComponent,
                    },
                ]
            },
            {
                path: "account",
                component: AccountPageComponent,
            },
            {
                path: "settings",
                component: SettingsPageComponent,
            },
        ]
    }
];
