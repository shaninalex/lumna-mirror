import { Routes } from "@angular/router";
import { ForYou } from "./for-you/for-you";
import { provideWorkspaceFeature } from "@entities/workspace";
import { provideTaskFeature } from "@entities/task";
import { provideListFeature } from "@entities/list";
import { provideProjectFeature } from "@entities/project";

export const routes: Routes = [
    {
        path: "",
        providers: [
            provideWorkspaceFeature(),
            provideProjectFeature(),
            provideListFeature(),
            provideTaskFeature()
        ],
        children: [
            {
                path: "workspaces",
                loadChildren: () =>
                    import("@pages/workspaces/routes").then(
                        (m) => m.workspacesListRoutes
                    )
            },
            {
                path: "",
                children: [
                    {
                        path: "",
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
            }
        ]
    }
];
