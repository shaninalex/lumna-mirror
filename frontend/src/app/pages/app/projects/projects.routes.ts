import { Routes } from '@angular/router';

import { ProjectList } from './project-list/project-list';
import { ProjectDetail } from './project-detail/project-detail';
import { ProjectEditPage } from './project-edit/project-edit.page';
import { ProjectsContainer } from './container';

export const projectsRoutes: Routes = [
    {
        path: 'projects',
        component: ProjectsContainer,
        children: [
            {
                path: '',
                component: ProjectList,
            },
            {
                path: ':id',
                component: ProjectDetail,
            },
            {
                path: ':id/edit',
                component: ProjectEditPage,
            },
        ],
    },
];
