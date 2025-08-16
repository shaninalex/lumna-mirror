import {Routes} from '@angular/router';
import {ProjectListPageComponent} from './list/project-list-page.component';
import {ProjectDetailPageComponent} from './detail/project-detail-page.component';

export const projectRoutes: Routes = [
    {
        path: 'projects',
        component: ProjectListPageComponent,
    },
    {
        path: 'projects/:id',
        component: ProjectDetailPageComponent,
    },
];
