import {
    HttpRequest,
    HttpHandlerFn,
    HttpEvent,
    HttpClient,
    HttpErrorResponse,
} from '@angular/common/http';
import { inject } from '@angular/core';
import { BehaviorSubject, catchError, filter, Observable, switchMap, take, throwError } from 'rxjs';

let isRefreshing = false;
const refreshSubject = new BehaviorSubject<boolean | null>(null);

export function apiInterceptor(
    req: HttpRequest<unknown>,
    next: HttpHandlerFn,
): Observable<HttpEvent<unknown>> {
    const http = inject(HttpClient);
    const authReq = req.clone({
        withCredentials: true,
    });

    return next(authReq).pipe(
        catchError((error: HttpErrorResponse) => {
            if (error.status === 401 && !authReq.url.includes('/api/v1/auth/refresh')) {
                return handle401(authReq, next, http);
            }

            // TODO: handle other errors:
            // - 403
            // - 500
            // - 400

            return throwError(() => error);
        }),
    );
}

function handle401(
    req: HttpRequest<any>,
    next: HttpHandlerFn,
    http: HttpClient,
): Observable<HttpEvent<any>> {
    if (!isRefreshing) {
        isRefreshing = true;
        refreshSubject.next(null);

        return http.get('/api/v1/auth/refresh', { withCredentials: true }).pipe(
            switchMap(() => {
                isRefreshing = false;
                refreshSubject.next(true);
                return next(
                    req.clone({
                        withCredentials: true,
                    }),
                );
            }),
            catchError((err) => {
                isRefreshing = false;
                refreshSubject.next(false);
                return throwError(() => err);
            }),
        );
    }

    return refreshSubject.pipe(
        filter((result) => result === true),
        take(1),
        switchMap(() =>
            next(
                req.clone({
                    withCredentials: true,
                }),
            ),
        ),
    );
}
