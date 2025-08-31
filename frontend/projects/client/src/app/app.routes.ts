import { Routes } from '@angular/router';
import {authRoutes, mainRoutes, systemRoutes} from '../pages';

export const routes: Routes = [
    ...authRoutes,
    ...mainRoutes,
    ...systemRoutes,
];
