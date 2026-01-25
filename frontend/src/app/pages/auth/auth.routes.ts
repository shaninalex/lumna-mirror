import { Routes } from '@angular/router';
import { Login } from './login/login';

export const authRoutes: Routes = [
    { path: 'auth/login', component: Login, title: 'Login' },
];
