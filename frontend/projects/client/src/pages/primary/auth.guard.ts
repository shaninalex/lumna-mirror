import {inject} from '@angular/core';
import {CanMatchFn, Router} from '@angular/router';
import {Store} from '@ngrx/store';
import {AppState} from '@client/shared/store';
import {GetUserAction, selectUserState} from '@client/entities/user';
import {map, filter, take, tap} from 'rxjs';

export const authGuard: CanMatchFn = () => {
    const store = inject(Store<AppState>);
    const router = inject(Router);

    return store.select(selectUserState).pipe(
        tap(state => {
            if (!state.loaded) {
                store.dispatch(GetUserAction());
            }
        }),
        filter(state => state.loaded),
        take(1),
        map(state => {
            if (state.user) {
                return true;
            } else {
                return router.createUrlTree(['/auth/login']);
            }
        })
    );
};
