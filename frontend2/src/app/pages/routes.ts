import { Routes } from "@angular/router";

import { authCanActivate } from "@core";

export const routes: Routes = [
    {
        path: "auth",
        loadChildren: () => import("./auth/routes").then((m) => m.routes)
    },

    {
        path: "",
        pathMatch: "full",
        redirectTo: "/app"
    },
    {
        path: "app",
        canMatch: [authCanActivate],
        loadChildren: () => import("./app/routes").then((m) => m.routes)
    },

    {
        path: "",
        loadChildren: () => import("./static/routes").then((m) => m.routes)
    },

    {
        path: "",
        loadChildren: () => import("./system/routes").then((m) => m.routes)
    }
];
