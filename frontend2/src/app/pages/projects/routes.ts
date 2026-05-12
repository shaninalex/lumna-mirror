import { Routes } from "@angular/router";
import { ProjectWrapper } from "./wrapper";
import { BoardPage, SummaryPage, BacklogPage, ProjectListPage } from "./pages/";

export const routes: Routes = [
    {
        path: "project/:slug",
        component: ProjectWrapper,
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
    },
    {
        path: "projects",
        component: ProjectListPage
    }
];
