import { Routes } from '@angular/router';
import { InboxPage } from './inbox';
import { ProjectsPage } from './projects';
import { WorkspacesPage } from './workspaces';
import { BoardsPage } from './boards';
import { BoardPage } from './board';
import { BacklogPage } from './backlog/backlog.page';
import { TaskPage } from './task';
import { WorkspaceCreateComponent } from './workspace-create';
import { activeWorkspaceGuard, workspaceRedirectGuard } from './workspace.guard';
import { ProjectsCreatePage } from './projects-create';

export const routes: Routes = [
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
                path: 'inbox',
                component: InboxPage,
            },
            {
                path: 'projects',
                component: ProjectsPage,
            },
            {
                path: 'projects/create',
                component: ProjectsCreatePage,
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
                path: '',
                pathMatch: 'full',
                redirectTo: 'inbox',
            },
        ],
    },
    {
        path: '',
        pathMatch: 'full',
        canActivate: [workspaceRedirectGuard],
        children: [],
    },
];

