import { Routes } from '@angular/router';
import { TaskContainer } from './container';
import { TaskDetailComponent } from './task-detail';

export const taskRoutes: Routes = [
    {
        path: 'task',
        component: TaskContainer,
        children: [
            {
                path: ':id',
                component: TaskDetailComponent,
            },
        ],
    },
];
