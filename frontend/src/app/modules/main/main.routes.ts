import type { Routes } from '@angular/router';
import { 
    WorkspacesPage, 
    WorkspaceCreateComponent, 
    WorkspaceEntryPage, 
    InboxPage, 
    BoardsPage, 
    BoardPage, 
    BacklogPage, 
    TaskPage, 
    ProjectsPage, 
    ProjectsCreatePage,
    TaskCreatePage,
    TaskDetailPage,
} from '@pages';
import { authGuard } from './auth.guard';
import { activeProjectGuard } from './project.guard';
import { activeWorkspaceGuard } from './workspace.guard';
import { lastRouteRedirect } from './lastRouteRedirect';

export const routes: Routes = [
    {
        path: '',
        canMatch: [authGuard],
        children: [
            {
                path: 'workspaces',
                component: WorkspacesPage,
            },
            {
                path: 'workspaces/create',
                component: WorkspaceCreateComponent,
            },
            {
                path: 'w/:workspaceId',
                canActivate: [activeWorkspaceGuard],
                children: [
                    {
                        path: '',
                        component: WorkspaceEntryPage,
                    },
                    {
                        path: 'p/:projectId',
                        canActivate: [activeProjectGuard],
                        children: [
                            {
                                path: 'inbox',
                                component: InboxPage,
                            },
                            {
                                path: 'boards',
                                component: BoardsPage,
                            },
                            {
                                path: 'board',
                                component: BoardPage,
                            },
                            {
                                path: 'backlog',
                                component: BacklogPage,
                            },
                            {
                                path: 'task',
                                component: TaskPage,
                            },
                            {
                                path: 'task/create',
                                component: TaskCreatePage,
                            },
                            {
                                path: 'task/:id',
                                component: TaskDetailPage,
                            },
                            {
                                path: '',
                                pathMatch: 'full',
                                redirectTo: 'inbox',
                            },
                        ],
                    },
                    {
                        path: 'projects',
                        component: ProjectsPage,
                    },
                    {
                        path: 'projects/create',
                        component: ProjectsCreatePage,
                    },
                ],
            },
            {
                path: '',
                pathMatch: 'full',
                canActivate: [lastRouteRedirect],
                children: [],
            },
        ],
    },
];
