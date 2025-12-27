import {CanActivateFn, Router} from '@angular/router';
import {inject} from '@angular/core';
import {SessionStore} from '@core/store/session.store';

export const guestGuard: CanActivateFn = () => {
    const sessionStore = inject(SessionStore);
    const router = inject(Router)
    const status = sessionStore.status();

    if (status === 'authenticated') {
        return router.navigate(['/'])
    }

    return true
};
