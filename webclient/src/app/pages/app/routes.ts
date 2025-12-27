import {Routes} from "@angular/router";

import {Home} from "./home/home";
import {Projects} from './projects/projects';

import {authGuard} from './auth.guard';
import {projectsRoutes} from './projects/projects.routes';
import {boardRoutes} from './boards/boards.routes';

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
            ...projectsRoutes,
            ...boardRoutes,
        ]
    }
]
