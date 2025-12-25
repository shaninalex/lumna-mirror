import {ApplicationConfig, isDevMode, provideBrowserGlobalErrorListeners} from '@angular/core';
import {provideRouter} from '@angular/router';

import {routes} from './app.routes';
import {provideStoreDevtools} from '@ngrx/store-devtools';
import {provideHttpClient, withInterceptorsFromDi} from '@angular/common/http';
// import {GlobalInterceptor} from '@core/global.interceptor';

export const appConfig: ApplicationConfig = {
    providers: [
        provideBrowserGlobalErrorListeners(),
        provideHttpClient(withInterceptorsFromDi()),
        provideRouter(routes),
        provideStoreDevtools({ maxAge: 25, logOnly: !isDevMode() }),
        // {
        //     provide: HTTP_INTERCEPTORS,
        //     useClass: GlobalInterceptor,
        //     multi: true,
        // }
    ]
};
