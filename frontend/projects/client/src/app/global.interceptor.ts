import {inject, Injectable} from '@angular/core';
import {HttpErrorResponse, HttpEvent, HttpEventType, HttpHandler, HttpRequest} from '@angular/common/http';
import {catchError, Observable, tap, throwError} from 'rxjs';
import {APIResponse} from '@client/shared/models';
import {Store} from '@ngrx/store';
import {AppState} from '@client/shared/store';
import {AppErrorAction} from '@client/shared/store/app.actions';
import {GenericError} from '@ory/kratos-client';
import {Router} from '@angular/router';

@Injectable({providedIn: 'root'})
export class GlobalInterceptor {
    private store = inject(Store<AppState>);
    private router = inject(Router);

    intercept(req: HttpRequest<any>, handler: HttpHandler): Observable<HttpEvent<any>> {
        return handler.handle(req).pipe(
            tap(event => {
                if (event.type === HttpEventType.Response) {
                    // console.log(req.url, 'returned a response with status', event.status);
                }
            }),
            catchError((err: HttpErrorResponse) => {
                if ("error" in err.error) {
                    if ("id" in err.error.error && "code" in err.error.error) {
                        const kratosError: GenericError = err.error.error;
                        if (kratosError.id === "session_already_available") {
                            this.router.navigate(['/'])
                        }
                    }
                }

                if ("errors" in err.error) { // <= ???
                    this.handleResponseError(err.error as APIResponse<any>)
                }
                return throwError(() => err.error);
            })
        );
    }

    handleResponseError(err: APIResponse<any>): void {
        for (let e of err.errors) {
            this.store.dispatch(AppErrorAction({err: e}))
        }
    }
}
