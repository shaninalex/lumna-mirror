import { Routes } from "@angular/router";
import { WorkspaceWrapper } from "./wrapper";
import { WorkspaceSummary } from "./summary";
import { BoardPage } from "./board-page";
import { BacklogPage } from "./backlog-page";

export const routes: Routes = [
    {
        path: "workspace/:id",
        component: WorkspaceWrapper,
        children: [
            {
                path: "",
                pathMatch: "full",
                redirectTo: "summary"
            },
            {
                path: "summary",
                component: WorkspaceSummary
            },
            {
                path: "board",
                component: BoardPage
            },
            {
                path: "backlog",
                component: BacklogPage
            }
        ]
    }
];
