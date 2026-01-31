import { Routes } from '@angular/router';

import { Home } from './home/home';

import { projectsRoutes } from './projects/projects.routes';
import { boardRoutes } from './boards/boards.routes';
import { authGuard } from './guards/auth.guard';
import { Calendar } from './calendar/calendar';
import { ProjectsContainer } from './projects/container';

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
                component: ProjectsContainer,
                children: [...projectsRoutes, ...boardRoutes],
            },
            {
                path: 'calendar',
                component: Calendar,
            },
        ],
    },
];
