import { Routes } from '@angular/router';
import { authGuard } from './main';

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
    }
];
