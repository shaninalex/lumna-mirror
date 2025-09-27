import {inject, Injectable} from '@angular/core';
import {HttpErrorResponse, HttpHandler, HttpRequest} from '@angular/common/http';
import {BehaviorSubject, catchError, EMPTY, filter, switchMap, take, throwError} from 'rxjs';
import {Router} from '@angular/router';
import {AuthService} from '@client/entities/auth';

@Injectable({providedIn: 'root'})
export class GlobalInterceptor {
    private router = inject(Router);
    private authService = inject(AuthService);
    private isRefreshing = false;
    private refreshSubject = new BehaviorSubject<boolean>(true);

    intercept(req: HttpRequest<any>, next: HttpHandler) {
        return next.handle(req).pipe(
            catchError((err: HttpErrorResponse) => {
                if (err.status === 401) {
                    return this.handle401(req, next);
                }
                return throwError(() => err);
            })
        );
    }

    private handle401(req: HttpRequest<any>, next: HttpHandler) {
        if (!this.isRefreshing) {
            this.isRefreshing = true;
            this.refreshSubject.next(false);
            console.log("step 1")

            return this.authService.refresh().pipe(
                switchMap(() => {
                    console.log("step 2")
                    this.isRefreshing = false;
                    this.refreshSubject.next(true);
                    return next.handle(req);
                }),
                catchError(err => {
                    console.log("step 3")
                    this.isRefreshing = false;
                    this.refreshSubject.next(false); // unblock waiting requests
                    this.router.navigate(['/auth/login']);
                    return EMPTY;
                })
            );
        } else {
            console.log("step 4")
            return this.refreshSubject.pipe(
                take(1),
                switchMap(success => {
                    console.log("step 5", success)
                    if (success) {
                        return next.handle(req);
                    } else {
                        this.router.navigate(['/auth/login']);
                        return EMPTY;
                    }
                })
            );
        }
    }
}
