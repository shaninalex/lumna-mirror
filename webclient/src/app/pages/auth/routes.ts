import { Routes } from "@angular/router";
import { Login } from "./login/login";
import { Confirm } from "./confirm/confirm";
import {Restore} from './restore/restore';

export const routes: Routes = [
    {
        path: "auth/login",
        component: Login,
    },
    {
        path: "auth/confirm",
        component: Confirm,
    },
    {
        path: "auth/restore",
        component: Restore,
    }
]
