import { Routes } from '@angular/router';
import {authRoutes, mainRoutes} from '../pages';

export const routes: Routes = [
    ...authRoutes,
    ...mainRoutes,
];
