import { Routes } from '@angular/router';
import {mainRoutes, systemRoutes, authRoutes} from '../pages';

export const routes: Routes = [
    ...authRoutes,
    ...mainRoutes,
    ...systemRoutes,
];
