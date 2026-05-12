import { Routes } from "@angular/router";
import {
    WorkspaceDetailSummary,
    WorkspaceDetailWrapper,
    WorkspaceList
} from "./pages";
import { routes as projectRoutes } from "@pages/projects";

export const routes: Routes = [
    {
        path: ":workspace-slug",
        component: WorkspaceDetailWrapper,
        children: [
            {
                path: "",
                component: WorkspaceDetailSummary
            },
            ...projectRoutes
        ]
    }
];

export const routesWorkspaceRoot = [
    {
        path: "workspaces",
        component: WorkspaceList
    }
];
