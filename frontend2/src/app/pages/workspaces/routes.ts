import { Routes } from "@angular/router";
import { WorkspaceDetailSummary, WorkspaceDetailWrapper } from "./pages";

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
