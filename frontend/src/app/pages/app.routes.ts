import { Routes } from '@angular/router';
import { authGuard } from './main';
import { Page404 } from './system';

export const routes: Routes = [
    {
        path: "auth",
        loadChildren: () => import('./auth/auth.routes').then((m) => m.routes)
    },
    {
        path: "app",
        canMatch: [authGuard],
        loadChildren: () => import('./main/main.routes').then((m) => m.routes)
    },
    {
        path: '',
        pathMatch: 'full',
        redirectTo: '/app'
    },
    {
        path: "404",
        loadComponent: () => import('./system/page-404').then((m) => m.Page404)
    },
    {
        path: '**',
        loadComponent: () => import('./system/page-404').then((m) => m.Page404)
    }
];
