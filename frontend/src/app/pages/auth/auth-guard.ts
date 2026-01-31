import { inject } from '@angular/core';
import { CanActivateFn } from '@angular/router';
import { selectUser, UserState } from '@entities/user';
import { Store } from '@ngrx/store';
import { tap } from 'rxjs';

/**
 * Redirect authenticated users to home page
 * if they trying to access auth page
 *
 * @returns boolean
 */
export const authGuard: CanActivateFn = (route, state) => {
    const store = inject(Store<UserState>);
    store.select(selectUser).pipe(
        tap((user) => {
            console.log(user);
        }),
    );
    return true;
};
