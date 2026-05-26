import { Routes } from "@angular/router";
import {
    WorkspaceDetailSummary,
    WorkspaceDetailWrapper,
    WorkspaceList
} from "./pages";
import { WorkspaceCreatePage } from "./pages/workspace-create-page";

export const routes: Routes = [
    {
        path: ":workspace-slug",
        component: WorkspaceDetailWrapper,
        children: [
            {
                path: "",
                component: WorkspaceDetailSummary
            },
            {
                path: "",
                loadChildren: () =>
                    import("@pages/projects").then((m) => m.routes)
            }
        ]
    }
];

export const workspacesListRoutes: Routes = [
    {
        path: "",
        component: WorkspaceList
    },
    {
        path: "create",
        component: WorkspaceCreatePage
    }
];
