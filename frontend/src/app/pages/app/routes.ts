import { Routes } from '@angular/router';

import { Home } from './home/home';

import { authGuard } from './guards/auth.guard';
import { Calendar } from './calendar/calendar';
import { DashboardContainer } from './container';
import { boardRoutes } from './boards/boards.routes';
import { projectsRoutes } from './projects/projects.routes';
import { taskRoutes } from './task';

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
            ...projectsRoutes,
            ...boardRoutes,
            ...taskRoutes,
        ],
    },
];
