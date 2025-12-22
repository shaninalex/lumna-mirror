import { Routes } from "@angular/router";
import { Home } from "./home/home";
import {authGuard} from '@root/src/app/pages/app/auth.guard';

export const routes: Routes = [
    {
        path: "",
        component: Home,
        canMatch: [authGuard]
    }
]
