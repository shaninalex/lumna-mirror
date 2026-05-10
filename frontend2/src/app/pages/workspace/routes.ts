import { Routes } from "@angular/router";
import { WorkspaceWrapper } from "./wrapper";
import { BoardPage } from "./board-page";
import { BacklogPage } from "./backlog-page";
import { SummaryPage } from "./summary-page";

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
                component: SummaryPage
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
