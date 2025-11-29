import { ApplicationConfig, isDevMode, provideBrowserGlobalErrorListeners, provideZoneChangeDetection } from "@angular/core"
import { provideRouter, withRouterConfig } from "@angular/router"

import { routes } from "./app.routes"
import { provideStore } from "@ngrx/store"
import { provideStoreDevtools } from "@ngrx/store-devtools"
import { provideEffects } from "@ngrx/effects"
import { provideRouterStore } from "@ngrx/router-store"
import { HTTP_INTERCEPTORS, provideHttpClient, withInterceptors, withInterceptorsFromDi } from "@angular/common/http"
import { effects, reducers } from "./app.store"
import { GlobalInterceptor } from "@client/app/global.interceptor"
import { authInterceptor } from "@client/app/auth-request.interceptor"

export const appConfig: ApplicationConfig = {
	providers: [
		provideBrowserGlobalErrorListeners(),
		provideZoneChangeDetection({ eventCoalescing: true }),
		provideRouter(routes, withRouterConfig({ paramsInheritanceStrategy: "always" })),
		provideHttpClient(withInterceptorsFromDi(), withInterceptors([authInterceptor])),
		provideRouterStore(),
		provideStore(reducers),
		provideEffects(effects),
		provideStoreDevtools({ maxAge: 25, logOnly: !isDevMode() }),
		{
			provide: HTTP_INTERCEPTORS,
			useClass: GlobalInterceptor,
			multi: true,
		},
	],
}
