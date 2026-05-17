import { Routes } from "@angular/router";
import { AuthWrapper } from "./wrapper";
import { LoginPage } from "./pages";

export const routes: Routes = [
    {
        path: "",
        component: AuthWrapper,
        children: [
            {
                path: "login",
                component: LoginPage
            }
        ]
    }
];
