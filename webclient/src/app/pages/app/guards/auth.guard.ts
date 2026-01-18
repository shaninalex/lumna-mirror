import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';
import { sessionEvents } from '@core/store/session.store';
import { UserService, UserStore } from '@entities/user';
import { Dispatcher } from '@ngrx/signals/events';
import { catchError, map, of } from 'rxjs';

export const authGuard: CanActivateFn = () => {
    const userService = inject(UserService);
    const userStore = inject(UserStore);
    const router = inject(Router);
    const dispatch = inject(Dispatcher);

    return userService.getMe().pipe(
        map((user) => {
            userStore.setUser(user);
            dispatch.dispatch(sessionEvents.authenticated());
            return true;
        }),
        catchError(() => {
            router.navigate(['/auth/login']);
            return of(false);
        }),
    );
};
