import { Routes } from "@angular/router";
import { Page404 } from "./page-404/page-404";

export const routes: Routes = [
    {
        path: "404",
        component: Page404
    },
    {
        path: "**",
        redirectTo: "/404"
    }
];
