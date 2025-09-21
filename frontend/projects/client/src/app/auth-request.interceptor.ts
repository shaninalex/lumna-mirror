import {HttpHandlerFn, HttpRequest} from '@angular/common/http';
import {inject} from '@angular/core';
import {TokenService} from '@client/shared/common';

export function authInterceptor(req: HttpRequest<unknown>, next: HttpHandlerFn) {
    const authToken = inject(TokenService).getAuthToken();
    const newReq = req.clone({
        headers: req.headers.append('Authorization', `Bearer ${authToken}`),
    });
    return next(newReq);
}
