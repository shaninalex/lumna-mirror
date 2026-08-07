import { Routes } from '@angular/router';
import { InboxPage } from './inbox';
import { ProjectsPage } from './projects';
import { WorkspacesPage } from './workspaces';


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
        path: '',
        pathMatch: 'full',
        redirectTo: '/inbox'
    }
];
