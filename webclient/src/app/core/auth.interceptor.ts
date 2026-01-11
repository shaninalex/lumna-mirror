import {HttpErrorResponse, HttpHandlerFn, HttpInterceptorFn, HttpRequest} from '@angular/common/http';
import {inject} from '@angular/core';
import {BehaviorSubject, catchError, filter, switchMap, take, throwError} from 'rxjs';
import {LoginService} from '@core/api/login.service';
import {Dispatcher} from '@ngrx/signals/events';
import {sessionEvents} from '@core/store/session.store';

let isRefreshing = false;
const refreshSubject = new BehaviorSubject<boolean>(false);

export const authInterceptor: HttpInterceptorFn = (req: HttpRequest<unknown>, next: HttpHandlerFn) => {
    const loginService = inject(LoginService);
    const dispatcher = inject(Dispatcher);

    return next(req).pipe(
        catchError((error: HttpErrorResponse) => {
            if (error.status === 401 && !isRefreshEndpoint(req.url)) {
                return handleUnauthorized(req, next, loginService, dispatcher);
            }
            return throwError(() => error);
        })
    );
};

function isRefreshEndpoint(url: string): boolean {
    return url.includes('/auth/refresh');
}

function handleUnauthorized(
    req: HttpRequest<unknown>,
    next: HttpHandlerFn,
    loginService: LoginService,
    dispatcher: Dispatcher
) {
    if (!isRefreshing) {
        isRefreshing = true;
        refreshSubject.next(false);

        return loginService.Refresh().pipe(
            switchMap(() => {
                isRefreshing = false;
                refreshSubject.next(true);
                dispatcher.dispatch(sessionEvents.refreshSucceeded());
                return next(req);
            }),
            catchError((refreshError) => {
                isRefreshing = false;
                refreshSubject.next(false);
                dispatcher.dispatch(sessionEvents.sessionFailed(refreshError));
                return throwError(() => refreshError);
            })
        );
    }

    // Wait for ongoing refresh to complete, then retry
    return refreshSubject.pipe(
        filter(refreshed => refreshed),
        take(1),
        switchMap(() => next(req))
    );
}
