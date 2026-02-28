import {Routes} from '@angular/router';


import {authGuard} from './guards/auth.guard';
import {DashboardContainer} from './container';
import {Home} from '@pages/app/home';
import {Calendar} from '@pages/app/calendar';
import {SettingsPage, SettingsContainer} from '@pages/app/settings';
import {TaskDetailComponent, taskResolver} from '@pages/app/task';
import {BoardContainer, BoardEditPage, BoardPage, boardResolver} from '@pages/app/boards';
import {
    ProjectContainer,
    ProjectDetail,
    ProjectEditPage,
    ProjectList,
    projectResolver,
    ProjectsContainer,
} from '@pages/app/projects';

export const routes: Routes = [
    {
        path: '',
        canActivate: [authGuard],
        component: DashboardContainer,
        children: [
            {
                path: '',
                component: Home,
                title: 'Overview',
            },
            {
                path: 'calendar',
                component: Calendar,
            },
            {
                path: 'settings',
                component: SettingsContainer,
                children: [
                    {
                        path: '',
                        component: SettingsPage,
                        data: {title: 'Settings page'},
                    },
                ],
            },
            {
                path: 'projects',
                component: ProjectsContainer,
                children: [
                    {
                        path: '',
                        component: ProjectList,
                    },
                    {
                        path: ':id',
                        component: ProjectContainer,
                        children: [
                            {
                                path: '',
                                component: ProjectDetail,
                                resolve: {
                                    project: projectResolver,
                                },
                            },
                            {
                                path: 'edit',
                                component: ProjectEditPage,
                                resolve: {
                                    project: projectResolver,
                                },
                            },
                            {
                                path: 'boards/:id',
                                component: BoardContainer,
                                children: [
                                    {
                                        path: '',
                                        component: BoardPage,
                                        resolve: {
                                            board: boardResolver,
                                        },
                                    },
                                    {
                                        path: 'edit',
                                        component: BoardEditPage,
                                        resolve: {
                                            board: boardResolver,
                                        },
                                    },
                                    {
                                        path: 'task/:id',
                                        component: TaskDetailComponent,
                                        resolve: {
                                            task: taskResolver,
                                        },
                                    }
                                ],
                            },
                        ],
                    },
                ],
            },
        ],
    },
];
