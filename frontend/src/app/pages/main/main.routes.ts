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
import { activeProjectGuard } from './project.guard';
import { ProjectsCreatePage } from './projects-create';
import { WorkspaceEntryPage } from '@pages/main/workspace-entry/workspace-entry.page';

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
        canActivate: [workspaceRedirectGuard],
        children: [],
    },
];
