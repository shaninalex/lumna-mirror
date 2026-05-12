import { Routes } from "@angular/router";
import { ForYou } from "./for-you/for-you";
import { routes as workspacesRoutes } from "@pages/workspaces";
import { ApplicationWrapper } from "./wrapper";

export const routes: Routes = [
    {
        path: "app",
        component: ApplicationWrapper,
        children: [
            {
                path: "",
                pathMatch: "full",
                component: ForYou
            },
            {
                path: "for-you",
                component: ForYou
            },
            ...workspacesRoutes
        ]
    }
];
