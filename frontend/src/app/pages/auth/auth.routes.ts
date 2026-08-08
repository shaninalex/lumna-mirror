import { Routes } from '@angular/router';
import { AuthRoot } from '@pages/auth/auth.root';
import { LoginPage } from '@pages/auth/login/login.page';

export const routes: Routes = [
    {
        path: '',
        component: AuthRoot,
        children: [
            {
                path: 'login',
                component: LoginPage,
            },
            {
                path: '',
                pathMatch: 'full',
                redirectTo: '/auth/login'
            }
        ]
    }
];
