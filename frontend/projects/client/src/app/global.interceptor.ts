import {inject, Injectable} from '@angular/core';
import {HttpErrorResponse, HttpEvent, HttpEventType, HttpRequest} from '@angular/common/http';
import {HttpHandler} from '@angular/common/http';
import {catchError, Observable, tap, throwError} from 'rxjs';
import {APIResponse} from '@client/shared/models';
import {ERRORS} from '@client/shared/common/errors';
import {Store} from '@ngrx/store';
import {AppState} from '@client/shared/store';
import {AppErrorAction} from '@client/shared/store/app.actions';

@Injectable({ providedIn: 'root' })
export class GlobalInterceptor {
    private store = inject(Store<AppState>);

    intercept(req: HttpRequest<any>, handler: HttpHandler): Observable<HttpEvent<any>> {
        return handler.handle(req).pipe(
            tap(event => {
                if (event.type === HttpEventType.Response) {
                    console.log(req.url, 'returned a response with status', event.status);
                }
            }),
            catchError((err: HttpErrorResponse) => {
                if ("errors" in err.error) {
                    this.handleResponseError(err.error as APIResponse<any>)
                }
                return throwError(() => err.error);
            })
        );
    }

    handleResponseError(err: APIResponse<any>): void {
        for (let e of err.errors) {
            this.store.dispatch(AppErrorAction({ err: e }))
        }
    }
}
