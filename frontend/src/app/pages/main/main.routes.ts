import { Routes } from '@angular/router';
import { InboxPage } from './inbox';
import { ProjectsPage } from './projects';
import { WorkspacesPage } from './workspaces';
import { BoardsPage } from './boards';
import { BoardPage } from './board';
import { BacklogPage } from './backlog/backlog.page';


export const routes: Routes = [
    {
        path: 'inbox',
        component: InboxPage,
    },
    {
        path: 'projects',
        component: ProjectsPage,
    },
    {
        path: 'workspaces',
        component: WorkspacesPage,
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
        path: '',
        pathMatch: 'full',
        redirectTo: '/inbox'
    }
];
