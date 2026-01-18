import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';
import { UserService, UserStore } from '@entities/user';
import { catchError, map, of } from 'rxjs';

export const authGuard: CanActivateFn = () => {
    const userService = inject(UserService);
    const userStore = inject(UserStore);
    const router = inject(Router);

    return userService.getMe().pipe(
        map((user) => {
            userStore.setUser(user);
            return true;
        }),
        catchError(() => {
            router.navigate(['/auth/login']);
            return of(false);
        }),
    );
};
