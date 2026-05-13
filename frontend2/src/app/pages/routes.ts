import { Routes } from "@angular/router";

import { routes as authRoutes } from "./auth";

export const routes: Routes = [
    ...authRoutes,

    {
        path: "",
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
