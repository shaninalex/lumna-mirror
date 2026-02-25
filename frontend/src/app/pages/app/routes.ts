import { Routes } from '@angular/router';

import { Home } from './home/home';

import { authGuard } from './guards/auth.guard';
import { Calendar } from './calendar/calendar';
import { DashboardContainer } from './container';
import { taskRoutes } from './task';
import { SettingsContainer } from '@pages/app/settings/settings-container';
import { SettingsPage } from '@pages/app/settings/settings-page/settings-page';
import { ProjectsContainer } from '@pages/app/projects/container';
import { ProjectList } from '@pages/app/projects/project-list/project-list';
import { ProjectContainer, ProjectDetail } from '@pages/app/projects/project-detail/project-detail';
import { projectResolver } from '@pages/app/projects/resolver';
import { ProjectEditPage } from '@pages/app/projects/project-edit/project-edit.page';
import { BoardContainer, BoardPage } from '@pages/app/boards/board-page/board-page';
import { boardResolver } from '@pages/app/boards/resolver';
import { BoardEditPage } from '@pages/app/boards/board-edit-page/board-edit-page';
import {TaskContainer} from '@pages/app/task/container';
import {taskResolver} from '@pages/app/task/task-resolver';
import {TaskDetailComponent} from '@pages/app/task/task-detail';

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
                        data: { title: 'Settings page' },
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
                                ],
                            },
                        ],
                    },
                ],
            },
            {
                path: 'task',
                component: TaskContainer,
                children: [
                    {
                        path: ':id',
                        component: TaskDetailComponent,
                        resolve: {
                            task: taskResolver
                        }
                    },
                ],
            },
        ],
    },
];
