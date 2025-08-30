import {Routes} from '@angular/router';
import {ProjectListPageComponent} from '@client/pages/main/projects/project-list/project-list-page.component';
import {ProjectDetailPageComponent} from '@client/pages/main/projects/project-detail/project-detail-page.component';
import {TaskDetailComponent} from '@client/pages/main/projects/task-detail/task-detail.component';

export const projectRoutes: Routes = [
    {
        path: 'projects',
        component: ProjectListPageComponent,
    },
    {
        path: 'projects/:projectCode',
        component: ProjectDetailPageComponent,
    },
    {
        path: 'projects/:projectKey/:taskCode',
        component: TaskDetailComponent,
    },
];
