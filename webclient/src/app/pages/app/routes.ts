import {Routes} from "@angular/router";
import {Home} from "./home/home";
import {Projects} from './projects/projects';
import {ProjectEdit} from './projects/project-edit/project-edit';

import {authGuard} from '@root/src/app/pages/app/auth.guard';
import { ProjectDetail } from "./projects/project-detail/project-detail";
import { ProjectList } from "./projects/project-list/project-list";
import {projectResolver} from '@root/src/app/pages/app/projects/project.resolver';

export const routes: Routes = [
    {
        path: "",
        component: Home,
        canMatch: [authGuard],
        title: "Overview"
    },
    {
        path: "projects",
        component: Projects,
        canMatch: [authGuard],
        children: [
            {
                path: "",
                component: ProjectList,
                canMatch: [authGuard],
            },
            {
                path: ":id",
                component: ProjectDetail,
                canMatch: [authGuard],
                resolve: {
                    project: projectResolver
                }
            },
            {
                path: ":id/edit",
                component: ProjectEdit,
                canMatch: [authGuard],
                resolve: {
                    project: projectResolver
                }
            }
        ]
    }
]
