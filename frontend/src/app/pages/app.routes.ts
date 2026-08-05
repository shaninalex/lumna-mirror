import { Routes } from '@angular/router';
import {routes as authRoutes } from './auth'
import {routes as mainRoutes } from './main'

export const routes: Routes = [
    ...authRoutes,
    ...mainRoutes,
];
