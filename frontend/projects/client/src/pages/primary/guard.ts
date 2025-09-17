import {map, Observable, of, switchMap} from 'rxjs';
import {inject, Injectable} from '@angular/core';
import {Store} from '@ngrx/store';
import {CanMatch, Route, Router, UrlSegment} from '@angular/router';
import {AuthService, selectSession, SetSessionAction} from '@client/entities/auth';
import {catchError, take} from 'rxjs/operators';
import {HttpErrorResponse} from '@angular/common/http';
import {AppState} from '@client/shared/store';
// import {GetUserAction} from '@client/entities/user/model/user.actions';

@Injectable({providedIn: 'root'})
export class CanMatchPrimarySection implements CanMatch {
    private api = inject(AuthService);
    private store = inject(Store<AppState>);
    private router = inject(Router);

    canMatch(route: Route, segments: UrlSegment[]): Observable<boolean> {
        return this.store.select(selectSession).pipe(
            take(1),
            switchMap(session => {
                if (session) return of(true)
                return this.api.session().pipe(
                    map(fetchedSession => {
                        if (fetchedSession) {
                            this.store.dispatch(SetSessionAction({session: fetchedSession}));
                            // this.store.dispatch(GetUserAction())
                            return true;
                        }
                        this.router.navigate(['/auth/login']);
                        return false;
                    }),
                    catchError((err: HttpErrorResponse) => {
                        let params: { [key: string]: any } = {}
                        if (err.error.error?.id === 'session_aal2_required' && err.error.redirect_browser_to) {
                            params["aal"] = "aal2"
                        }
                        this.router.navigate(['/auth/login'], {queryParams: params});
                        return of(false);
                    })
                )
            })
        )
    }
}
