import { Routes } from "@angular/router";
import { AuthWrapper } from "./wrapper";
import { AcceptInvitePage, LoginPage } from "./pages";

export const routes: Routes = [
    {
        path: "",
        component: AuthWrapper,
        children: [
            {
                path: "login",
                component: LoginPage
            },
            {
                path: "accept-invite",
                component: AcceptInvitePage
            }
        ]
    }
];
