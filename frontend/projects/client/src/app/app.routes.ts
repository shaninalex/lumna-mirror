import { Routes } from '@angular/router';
import {authRoutes} from '@client/pages/auth';
import {mainRoutes} from '@client/pages/main';

export const routes: Routes = [
    ...authRoutes,
    ...mainRoutes,
];
