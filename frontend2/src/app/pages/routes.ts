import { Routes } from "@angular/router";

import { routes as appRoutes } from "./app";
import { routes as authRoutes } from "./auth";
import { routes as staticRoutes } from "./static";
import { routes as systemRoutes } from "./system";

export const routes: Routes = [
    ...authRoutes,
    ...appRoutes,
    ...staticRoutes,
    ...systemRoutes
];
