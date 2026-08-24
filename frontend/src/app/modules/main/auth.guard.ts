import { inject } from '@angular/core';
import type { CanActivateFn} from '@angular/router';
import { Router } from '@angular/router';
import { actionSession } from '@core';
import { UserApi, actionUser } from '@entities/user';
import { Store } from '@ngrx/store';
import { catchError, map, of } from 'rxjs';

export const authGuard: CanActivateFn = () => {
    const userApi = inject(UserApi);
    const router = inject(Router);
    const store = inject(Store);

    return userApi.me().pipe(
        map((user) => {
            store.dispatch(actionUser.set({ user }));
            store.dispatch(actionSession.authenticated());
            return true;
        }),
        catchError(() => {
            router.navigate(['/auth/login']);
            return of(false);
        }),
    );
};
