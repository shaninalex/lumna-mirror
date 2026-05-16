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

import { reducers, effects, ApplicationInit, sessionInterceptor } from "@core";
import { provideStoreDevtools } from "@ngrx/store-devtools";
import { provideHttpClient, withInterceptors } from "@angular/common/http";

export const appConfig: ApplicationConfig = {
    providers: [
        provideBrowserGlobalErrorListeners(),
        provideHttpClient(withInterceptors([sessionInterceptor])),
        provideRouter(routes),
        provideStore(reducers),
        provideEffects(effects),
        // provideRouterStore(),
        provideStoreDevtools({ maxAge: 25, logOnly: !isDevMode() }),
        provideAppInitializer(ApplicationInit)
    ]
};
