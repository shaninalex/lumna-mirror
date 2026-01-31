import { ApplicationConfig, isDevMode, provideBrowserGlobalErrorListeners } from '@angular/core';
import { provideHttpClient, withInterceptors } from '@angular/common/http';
import { provideRouter } from '@angular/router';

import { provideEffects } from '@ngrx/effects';
import { provideStoreDevtools } from '@ngrx/store-devtools';
import { provideStore } from '@ngrx/store';

import { effects, reducers, refreshTokenInterceptor } from '@core';
import { routes as mainRoutes } from '@pages';

export const appConfig: ApplicationConfig = {
    providers: [
        provideBrowserGlobalErrorListeners(),
        provideHttpClient(withInterceptors([refreshTokenInterceptor])),
        provideRouter(mainRoutes),
        provideEffects(effects),
        provideStore(reducers),
        provideStoreDevtools({ maxAge: 25, logOnly: !isDevMode() }),
    ],
};
