import { Routes } from "@angular/router";
import { Page404 } from "./page-404/page-404";
import { OnboardingPage } from "./onboarding";

export const routes: Routes = [
    {
        path: "onboarding",
        component: OnboardingPage
    },
    {
        path: "404",
        component: Page404
    },
    {
        path: "**",
        redirectTo: "/404"
    }
];
