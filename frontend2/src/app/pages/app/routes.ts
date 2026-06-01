import { Routes } from "@angular/router";
import { provideWorkspaceFeature } from "@entities/workspace";
import { provideTaskFeature } from "@entities/task";
import { provideListFeature } from "@entities/list";
import { provideProjectFeature } from "@entities/project";
import { provideSprintFeature } from "@entities/sprint";

export const routes: Routes = [
    {
        path: "",
        providers: [
            provideWorkspaceFeature(),
            provideProjectFeature(),
            provideListFeature(),
            provideTaskFeature(),
            provideSprintFeature()
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
                        pathMatch: "full",
                        redirectTo: "/app/workspaces"
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
