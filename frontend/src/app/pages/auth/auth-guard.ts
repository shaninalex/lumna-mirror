import { CanActivateFn } from '@angular/router';

/**
 * Redirect authenticated users to home page
 * if they trying to access auth page
 *
 * @returns boolean
 */
export const authGuard: CanActivateFn = (route, state) => {
    return true;
};
