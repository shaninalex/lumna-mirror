import { Routes } from "@angular/router";
import { Login } from "./login/login";
import { Confirm } from "./confirm/confirm";
import {Restore} from './restore/restore';
import {guestGuard} from '@root/src/app/pages/guest.guard';

export const routes: Routes = [
    {
        path: "auth/login",
        component: Login,
        title: "Auth Login",
        canActivate: [guestGuard],
    },
    {
        path: "auth/confirm",
        component: Confirm,
        title: "Auth Confirm",
        canActivate: [guestGuard],
    },
    {
        path: "auth/restore",
        component: Restore,
        title: "Auth Restore",
        canActivate: [guestGuard],
    }
]
