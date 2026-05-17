import { Routes } from "@angular/router";
import { ForYou } from "./for-you/for-you";
import { ApplicationWrapper } from "./wrapper";
import { provideWorkspaceFeature } from "@entities/workspace";

export const routes: Routes = [
    {
        path: "",
        // canActivate: [authBootstrapGuard],
        providers: [provideWorkspaceFeature()],
        children: [
            {
                path: "",
                pathMatch: "full",
                redirectTo: "/app"
            },
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
                    {
                        path: "",
                        loadChildren: () =>
                            import("@pages/workspaces/routes").then(
                                (m) => m.routes
                            )
                    }
                ]
            },
            {
                path: "workspaces",
                children: [
                    {
                        path: "",
                        loadComponent: () =>
                            import("@pages/workspaces/pages/workspace-list/workspace-list").then(
                                (m) => m.WorkspaceList
                            )
                    }
                ]
            }
        ]
    }
];
