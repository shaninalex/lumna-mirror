import { Routes } from "@angular/router";
import { ApplicationWrapper } from "@pages/app/wrapper";
import {
    WorkspaceArchivedPage,
    WorkspaceCreatePage,
    WorkspaceDetailSummary,
    WorkspaceList
} from "./pages";
import { WorkspacesWrapper } from "@pages/workspaces/workspaces-wrapper";

// routes goes under ApplicationWrapper ( <application-wrapper> )
export const routes: Routes = [
    {
        path: ":workspace-slug",
        component: ApplicationWrapper,
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

// routes goes under WorkspacesWrapper ( <workspaces-wrapper> )
export const workspacesListRoutes: Routes = [
    {
        path: "",
        component: WorkspacesWrapper,
        children: [
            {
                path: "",
                component: WorkspaceList
            },
            {
                path: "create",
                component: WorkspaceCreatePage
            },
            {
                path: "archived",
                component: WorkspaceArchivedPage
            }
        ]
    }
];
