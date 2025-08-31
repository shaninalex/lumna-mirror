import { Routes } from '@angular/router';
import {AuthRoot} from './auth-root';
import {Login} from './login/login';
import {Register} from './register/registration';
import {Verification} from './verification/verification';

export const authRoutes: Routes = [
    {
        path: "auth",
        component: AuthRoot,
        children: [
            {
                path: "login",
                component: Login,
            },
            {
                path: "registration",
                component: Register,
            },
            {
                path: "verification",
                component: Verification,
            }
        ]
    }
];
