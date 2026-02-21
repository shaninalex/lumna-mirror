import { Routes } from '@angular/router';
import { TaskContainer } from './container';
import { TaskDetailComponent } from './task-detail';
import { taskResolver } from '@pages/app/task/task-resolver';

export const taskRoutes: Routes = [
    {
        path: 'task',
        component: TaskContainer,
        children: [
            {
                path: ':id',
                component: TaskDetailComponent,
                resolve: {
                    task: taskResolver
                }
            },
        ],
    },
];
