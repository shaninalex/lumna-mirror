import {ApplicationConfig, isDevMode, provideBrowserGlobalErrorListeners} from '@angular/core';
import { provideRouter } from '@angular/router';

import { routes } from '@pages';
import {provideStore} from '@ngrx/store';
import {provideEffects} from '@ngrx/effects';
import { provideStoreDevtools } from '@ngrx/store-devtools';
import { effects, reducers } from './store';
import { provideHttpClient, withInterceptors } from '@angular/common/http';
import { apiInterceptor } from './api.interceptor';


export const appConfig: ApplicationConfig = {
    providers: [
        provideBrowserGlobalErrorListeners(),
        provideHttpClient(withInterceptors([apiInterceptor])),
        provideRouter(routes),
        provideEffects(effects),
        provideStore(reducers),
        provideStoreDevtools({ maxAge: 25, logOnly: !isDevMode() }),
    ]
};
