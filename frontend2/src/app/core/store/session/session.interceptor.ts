import {
    HttpRequest,
    HttpHandlerFn,
    HttpEvent,
    HttpClient,
    HttpErrorResponse,
    HttpInterceptorFn
} from "@angular/common/http";
import { inject } from "@angular/core";
import { Store } from "@ngrx/store";
import {
    BehaviorSubject,
    catchError,
    EMPTY,
    filter,
    Observable,
    switchMap,
    take,
    throwError
} from "rxjs";
import { actionSessionFailed, SessionApi } from ".";

export const sessionInterceptor: HttpInterceptorFn = (
    req: HttpRequest<unknown>,
    next: HttpHandlerFn
): Observable<HttpEvent<unknown>> => {
    const refreshService = inject(SessionApi);
    const store = inject(Store);

    const authReq = req.clone({
        withCredentials: true
    });

    return next(authReq).pipe(
        catchError((error: HttpErrorResponse) => {
            const is401 = error.status === 401;

            const isRefreshCall = authReq.url.includes("/api/v1/auth/refresh");

            if (!is401 || isRefreshCall) {
                return throwError(() => error);
            }

            return refreshService.refresh().pipe(
                switchMap(() => next(authReq)),

                catchError((refreshError) => {
                    store.dispatch(actionSessionFailed());

                    return throwError(() => refreshError);
                })
            );
        })
    );
};

// let isRefreshing = false;
// const refreshSubject = new BehaviorSubject<boolean | null>(null);

// export function refreshTokenInterceptor(
//     req: HttpRequest<unknown>,
//     next: HttpHandlerFn
// ): Observable<HttpEvent<unknown>> {
//     const http = inject(HttpClient);
//     const store = inject(Store);
//     const authReq = req.clone({
//         withCredentials: true
//     });

//     return next(authReq).pipe(
//         catchError((error: HttpErrorResponse) => {
//             if (
//                 error.status === 401 &&
//                 !authReq.url.includes("/api/v1/auth/refresh")
//             ) {
//                 return handle401(authReq, next, http).pipe(
//                     catchError(() => {
//                         store.dispatch(actionSessionFailed());
//                         return EMPTY;
//                     })
//                 );
//             }

//             return throwError(() => error);
//         })
//     );
// }

// function handle401(
//     req: HttpRequest<any>,
//     next: HttpHandlerFn,
//     http: HttpClient
// ): Observable<HttpEvent<any>> {
//     if (!isRefreshing) {
//         isRefreshing = true;
//         refreshSubject.next(null);

//         return http.get("/api/v1/auth/refresh", { withCredentials: true }).pipe(
//             switchMap(() => {
//                 isRefreshing = false;
//                 refreshSubject.next(true);
//                 return next(
//                     req.clone({
//                         withCredentials: true
//                     })
//                 );
//             }),
//             catchError((err) => {
//                 isRefreshing = false;
//                 refreshSubject.next(false);
//                 return throwError(() => err);
//             })
//         );
//     }

//     return refreshSubject.pipe(
//         filter((result) => result === true),
//         take(1),
//         switchMap(() =>
//             next(
//                 req.clone({
//                     withCredentials: true
//                 })
//             )
//         )
//     );
// }
