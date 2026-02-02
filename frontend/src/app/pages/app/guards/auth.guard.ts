import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';
import { actionSessionAuthenticated } from '@core/store/session.actions';
import { actionUserSet, UserApi } from '@entities/user';
import { Store } from '@ngrx/store';
import { catchError, map, of } from 'rxjs';

export const authGuard: CanActivateFn = () => {
    const userApi = inject(UserApi);
    const router = inject(Router);
    const store = inject(Store);

    return userApi.me().pipe(
        map((user) => {
            store.dispatch(actionUserSet({ user }));
            store.dispatch(actionSessionAuthenticated());
            return true;
        }),
        catchError(() => {
            router.navigate(['/auth/login']);
            return of(false);
        }),
    );
};
