import {CanMatchFn, Route, Router, UrlSegment,} from '@angular/router';
import { inject } from '@angular/core';
import { Store } from '@ngrx/store';
import { GetSessionAction, selectSession } from '../model';
import { AppState } from '@client/shared/store';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';


export const AuthGuard: CanMatchFn = (route: Route, segments: UrlSegment[]) => {
    for (let i = 0; i < segments.length; i++) {
        if (segments[i].path === 'auth') {
            return true
        }
    }

    const store = inject(Store<AppState>);
    const router = inject(Router);

    store.select(selectSession).pipe(
        takeUntilDestroyed()
    ).subscribe({
        next: session => {
            if (!session) {
                router.navigate(['/auth/login'])
                store.dispatch(GetSessionAction())
            }
            const expDate = new Date(session?.expires_at as string);
            const currentDate = new Date();

            // If it outdated - renew session
            if (expDate < currentDate) {
                // TODO: renew session
                // Create new action "RenewSessionAction()"
                store.dispatch(GetSessionAction())
            }
        },
        error: (err) => {
            console.error(err)
        }
    })

    return true;
};
