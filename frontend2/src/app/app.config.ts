import { ApplicationConfig, isDevMode, provideAppInitializer, provideBrowserGlobalErrorListeners } from "@angular/core";
import { provideRouter } from "@angular/router";

import { routes } from "@pages";
import { provideStore } from "@ngrx/store";
import { provideEffects } from "@ngrx/effects";

import { ApplicationInit, LumnaPrimeNGTheme, provideApplicationFeature, sessionInterceptor } from "@core";
import { provideStoreDevtools } from "@ngrx/store-devtools";
import { provideHttpClient, withInterceptors } from "@angular/common/http";
import { provideRouterStore } from "@ngrx/router-store";
import { provideUserFeature } from "@entities/user";
import { provideSessionFeature } from "@core/store/session/provider";

// ui lib
import { providePrimeNG } from "primeng/config";

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

        provideStoreDevtools({ maxAge: 25, logOnly: !isDevMode() }),
        provideAppInitializer(ApplicationInit),

        providePrimeNG({
            theme: {
                preset: LumnaPrimeNGTheme,
                options: {
                    darkModeSelector: ".dark-mode"
                }
            }
        })
    ]
};
