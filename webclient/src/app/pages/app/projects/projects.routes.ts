import {Routes} from '@angular/router';
import {projectResolver} from '../resolvers/project.resolver';

import {ProjectList} from './project-list/project-list';
import {ProjectDetail} from './project-detail/project-detail';
import {ProjectEdit} from './project-edit/project-edit';

export const projectsRoutes: Routes = [
    {
        path: "",
        component: ProjectList,
    },
    {
        path: ":id",
        component: ProjectDetail,
        resolve: {
            project: projectResolver
        }
    },
    {
        path: ":id/edit",
        component: ProjectEdit,
        resolve: {
            project: projectResolver
        }
    },
]
