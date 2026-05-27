import { Routes } from "@angular/router";
import { WorkspaceDetailSummary, WorkspaceList } from "./pages";
import { WorkspaceCreatePage } from "./pages/workspace-create-page";
import { ApplicationWrapper } from "@pages/app/wrapper";

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
