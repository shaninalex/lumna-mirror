import {inject} from '@angular/core';
import {CanMatchFn, Router} from '@angular/router';
import {Store} from '@ngrx/store';
import {AppState} from '@client/shared/store';
import {GetUserAction, selectUser} from '@client/entities/user';
import {map, filter, take, tap} from 'rxjs';

export const authGuard: CanMatchFn = () => {
    const store = inject(Store<AppState>);
    const router = inject(Router);

    return store.select(selectUser).pipe(
        tap(user => {
            if (user === undefined) {
                // 🔥 trigger backend call once
                store.dispatch(GetUserAction());
            }
        }),
        // wait until user is resolved (either object or null)
        filter(user => user !== undefined),
        take(1),
        map(user => {
            if (user) {
                return true; // ✅ allow
            } else {
                return router.createUrlTree(['/auth/login']); // ❌ redirect
            }
        })
    );
};
