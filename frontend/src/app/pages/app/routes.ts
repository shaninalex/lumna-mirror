import { Routes } from '@angular/router';

import { Home } from './home/home';
import { Projects } from './projects/projects';

import { projectsRoutes } from './projects/projects.routes';
import { boardRoutes } from './boards/boards.routes';
import { authGuard } from './guards/auth.guard';
import { Calendar } from './calendar/calendar';

export const routes: Routes = [
    {
        path: '',
        canActivate: [authGuard],
        children: [
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
            {
                path: 'calendar',
                component: Calendar,
            },
        ],
    },
];
