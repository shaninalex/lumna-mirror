import { Routes } from '@angular/router';
import { Login } from './login/login';
import { AuthContainer } from './container';
import { authGuard } from './auth-guard';

export const authRoutes: Routes = [
    {
        path: 'auth',
        component: AuthContainer,
        canActivate: [authGuard],
        children: [{ path: 'login', component: Login, title: 'Login' }],
    },
];
