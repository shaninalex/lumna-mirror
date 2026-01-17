import {
    HttpErrorResponse,
    HttpHandlerFn,
    HttpInterceptorFn,
    HttpRequest,
} from '@angular/common/http';
import { inject } from '@angular/core';
import { catchError, EMPTY, throwError } from 'rxjs';
import { Dispatcher } from '@ngrx/signals/events';
import { sessionEvents } from '@core/store/session.store';

export const authInterceptor: HttpInterceptorFn = (
    req: HttpRequest<unknown>,
    next: HttpHandlerFn,
) => {
    const dispatcher = inject(Dispatcher);
    return next(req).pipe(
        catchError((error: HttpErrorResponse) => {
            if (error.status === 401) {
                dispatcher.dispatch(sessionEvents.sessionFailed(error.error));
                return throwError(() => error);
            }
            return EMPTY;
        }),
    );
};
