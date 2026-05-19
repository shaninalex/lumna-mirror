import {
    ApplicationConfig,
    provideBrowserGlobalErrorListeners,
    isDevMode,
    provideAppInitializer
} from "@angular/core";
import { provideRouter } from "@angular/router";

import { routes } from "@pages";
import { provideStore } from "@ngrx/store";
import { provideEffects } from "@ngrx/effects";

import {
    ApplicationInit,
    sessionInterceptor,
    provideApplicationFeature
} from "@core";
import { provideStoreDevtools } from "@ngrx/store-devtools";
import { provideHttpClient, withInterceptors } from "@angular/common/http";
import { provideRouterStore } from "@ngrx/router-store";
import { provideUserFeature } from "@entities/user";
import { provideSessionFeature } from "@core/store/session/provider";
import { provideOnboardingFeature } from "@features/onboarding";

export const appConfig: ApplicationConfig = {
    providers: [
        provideBrowserGlobalErrorListeners(),
        provideHttpClient(withInterceptors([sessionInterceptor])),
        provideRouter(routes),
        provideStore(),
        provideEffects(),
        provideRouterStore(),

        // global features
        provideApplicationFeature(),
        provideSessionFeature(),
        provideUserFeature(),
        provideOnboardingFeature(),

        provideStoreDevtools({ maxAge: 25, logOnly: !isDevMode() }),
        provideAppInitializer(ApplicationInit)
    ]
};
