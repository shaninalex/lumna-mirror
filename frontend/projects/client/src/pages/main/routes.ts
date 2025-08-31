import { Routes } from '@angular/router';
import {MainRoot} from './main-root';
import {Overview} from './overview/overview';
import {AuthGuard} from '@client/entities/session';

export const mainRoutes: Routes = [
    {
        path: "",
        component: MainRoot,
        canMatch: [AuthGuard],
        children: [
            {
                path: "",
                component: Overview,
            }
        ]
    }
];
