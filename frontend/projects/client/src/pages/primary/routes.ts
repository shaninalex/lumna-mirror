import {ResolveFn, Routes} from '@angular/router';
import {PrimaryRoot} from './primary-root';
import {Overview} from './overview/overview';
import {overviewResolver} from './overview/overview.resolver';
import {ProjectsListPageComponent} from './projects-list/projects-list-page.component'
import {projectListResolver} from './projects-list/projects-list.resolver';
import {
    BoardViewPageComponent,
    ProjectDetailPageComponent,
    ProjectSettingsPageComponent,
    projectResolver,
} from '@client/pages/primary/project-detail';
import {ProjectsRootComponent} from '@client/pages/primary/projects-root';
import {SettingsPageComponent} from '@client/pages/primary/settings-page';
import {authGuard} from './auth.guard';

export const mainRoutes: Routes = [
    {
        path: "",
        component: PrimaryRoot,
        canMatch: [authGuard],
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
                        resolve: {project: projectResolver},
                        children: [
                            {
                                path: "",
                                component: BoardViewPageComponent,
                            },
                            {
                                path: "settings",
                                component: ProjectSettingsPageComponent,
                            },
                        ]
                    },
                    {
                        path: ":projectKey/settings",
                        component: ProjectDetailPageComponent,
                    },
                ]
            },
            {
                path: "settings",
                component: SettingsPageComponent,
            },
        ]
    }
];
