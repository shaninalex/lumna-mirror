import { Routes } from "@angular/router";
import { activeWorkspaceGuard, activeProjectGuard, workspaceRedirectGuard } from "@pages/main";
import { BacklogPage } from "@pages/main/backlog/backlog.page";
import { BoardPage } from "@pages/main/board";
import { BoardsPage } from "@pages/main/boards";
import { InboxPage } from "@pages/main/inbox";
import { ProjectsPage } from "@pages/main/projects";
import { ProjectsCreatePage } from "@pages/main/projects-create";
import { TaskPage } from "@pages/main/task";
import { WorkspaceCreateComponent } from "@pages/main/workspace-create";
import { WorkspaceEntryPage } from "@pages/main/workspace-entry";
import { WorkspacesPage } from "@pages/main/workspaces";


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
