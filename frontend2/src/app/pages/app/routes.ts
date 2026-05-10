import { Routes } from "@angular/router";
import { ForYou } from "./for-you/for-you";
import { routes as workspaceRouters } from "../workspace";
import { ApplicationWrapper } from "./wrapper";
import { WorkspacesList } from "./workspaces";

export const routes: Routes = [
    {
        path: "app",
        component: ApplicationWrapper,
        children: [
            {
                path: "",
                component: WorkspacesList
            },
            {
                path: "for-you",
                component: ForYou
            },
            ...workspaceRouters
        ]
    }
];
