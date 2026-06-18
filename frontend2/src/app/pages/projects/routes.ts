import { Routes } from "@angular/router";
import { ProjectWrapper } from "./wrapper";
import {
    BacklogPage,
    BoardPage,
    ProjectCreatePage,
    ProjectListPage,
    SummaryPage
} from "./pages/";

export const routes: Routes = [
    {
        path: "project/:project-id",
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
    },
    {
        path: "projects/create",
        component: ProjectCreatePage
    }
];
