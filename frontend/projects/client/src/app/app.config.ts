import {
    ApplicationConfig,
    isDevMode,
    provideBrowserGlobalErrorListeners,
    provideZoneChangeDetection
} from '@angular/core';
import {provideRouter} from '@angular/router';

import {routes} from './app.routes';
import {provideStore} from '@ngrx/store';
import {provideStoreDevtools} from '@ngrx/store-devtools';
import {provideEffects} from '@ngrx/effects';
import {provideRouterStore} from '@ngrx/router-store';
import {provideHttpClient, withInterceptorsFromDi} from '@angular/common/http';
import {reducers} from './app.store';

export const appConfig: ApplicationConfig = {
    providers: [
        provideBrowserGlobalErrorListeners(),
        provideZoneChangeDetection({eventCoalescing: true}),
        provideRouter(routes),
        provideEffects(),
        provideHttpClient(withInterceptorsFromDi()),
        provideRouterStore(),
        provideStore(reducers),
        provideStoreDevtools({maxAge: 25, logOnly: !isDevMode()}),
    ]
};
