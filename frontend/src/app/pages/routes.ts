import { Routes } from "@angular/router";
import { routes as appRoutes } from './app'
import { routes as systemRoutes } from './system'
import { routes as onboardingRoutes } from './onboarding'
import { authRoutes } from './auth'

export const routes: Routes = [
    ...appRoutes,
    ...authRoutes,
    ...onboardingRoutes,
    ...systemRoutes,
]
