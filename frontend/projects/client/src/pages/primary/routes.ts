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
                resolve: { overview: overviewResolver },
                data: { breadcrumb: "Home"},
            },
            {
                path: "projects",
                component: ProjectsRootComponent,
                resolve: { projects: projectListResolver },
                data: { breadcrumb: "Projects"},
                children: [
                    {
                        path: "",
                        component: ProjectsListPageComponent,
                    },
                    {
                        path: ":projectKey",
                        component: ProjectDetailPageComponent,
                        resolve: {project: projectResolver},
                        data: {
                            breadcrumb: (data: any, params: any) => data.project?.title ?? params['projectKey'] ?? 'Project'
                        },
                        children: [
                            {
                                path: "",
                                component: BoardViewPageComponent,
                                data: { breadcrumb: 'Board' },
                            },
                            {
                                path: "settings",
                                component: ProjectSettingsPageComponent,
                                data: { breadcrumb: "Settings"},
                            },
                        ]
                    }
                ]
            },
            {
                path: "settings",
                component: SettingsPageComponent,
                data: { breadcrumb: "Settings"},
            },
        ]
    }
];
