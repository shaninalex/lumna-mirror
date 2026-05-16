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
import { provideRouterStore } from "@ngrx/router-store";

import { reducers, effects, ApplicationInit } from "@core";
import { provideStoreDevtools } from "@ngrx/store-devtools";

export const appConfig: ApplicationConfig = {
    providers: [
        provideBrowserGlobalErrorListeners(),
        provideRouter(routes),
        provideStore(reducers),
        provideEffects(effects),
        // provideRouterStore(),
        provideStoreDevtools({ maxAge: 25, logOnly: !isDevMode() }),
        provideAppInitializer(ApplicationInit)
    ]
};
