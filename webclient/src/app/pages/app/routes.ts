import { Routes } from '@angular/router';

import { Home } from './home/home';
import { Projects } from './projects/projects';

import { projectsRoutes } from './projects/projects.routes';
import { boardRoutes } from './boards/boards.routes';

export const routes: Routes = [
    {
        path: '',
        component: Home,
        title: 'Overview',
    },
    {
        path: 'projects',
        component: Projects,
        children: [...projectsRoutes, ...boardRoutes],
    },
];
