import { Routes } from '@angular/router';
import {AuthRoot} from './auth-root';
import {Login} from './login/login';

export const authRoutes: Routes = [
    {
        path: "auth",
        component: AuthRoot,
        children: [
            {
                path: "login",
                component: Login,
            }
        ]
    }
];
