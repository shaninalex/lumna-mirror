import {ApplicationConfig, provideBrowserGlobalErrorListeners, provideZoneChangeDetection,} from "@angular/core";
import {provideRouter} from '@angular/router';

import {routes} from './app.routes';
import {provideStore} from '@ngrx/store';
import {provideEffects} from '@ngrx/effects';
import {provideHttpClient, withInterceptorsFromDi} from '@angular/common/http';

import {MAT_FORM_FIELD_DEFAULT_OPTIONS} from '@angular/material/form-field';


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
        provideEffects(),
        {
            provide: MAT_FORM_FIELD_DEFAULT_OPTIONS,
            useValue: {
                subscriptSizing: 'dynamic'
            }
        }
    ]
};
