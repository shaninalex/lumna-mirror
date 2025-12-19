import { Routes } from "@angular/router";
import { Login } from "./login/login";
import { Confirm } from "./confirm/confirm";

export const routes: Routes = [
    {
        path: "auth/login",
        component: Login,
    },
    {
        path: "auth/confirm",
        component: Confirm,
    }
]