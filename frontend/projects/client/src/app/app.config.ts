import {
    ApplicationConfig,
    provideBrowserGlobalErrorListeners,
    provideZoneChangeDetection,
} from "@angular/core";
import {provideRouter} from '@angular/router';

import {routes} from './app.routes';
import {provideStore} from '@ngrx/store';
import {provideEffects} from '@ngrx/effects';
import {provideHttpClient, withInterceptorsFromDi} from '@angular/common/http';
import {provideAnimationsAsync} from '@angular/platform-browser/animations/async';

export const appConfig: ApplicationConfig = {
    providers: [
        provideHttpClient(
            withInterceptorsFromDi(),
            // withInterceptors([toastInterceptor])
        ),
        provideBrowserGlobalErrorListeners(),
        provideZoneChangeDetection({eventCoalescing: true}),
        provideRouter(routes),
        provideStore(),
        provideEffects()
    ]
};
