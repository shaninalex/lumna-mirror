import {CanActivateFn, CanMatchFn, Router} from '@angular/router';
import {inject} from '@angular/core';
import {userEvents, UserStore} from '@entities/user';
import {Dispatcher} from '@ngrx/signals/events';
import {SessionStore} from '@core/store/session.store';
import {toObservable} from '@angular/core/rxjs-interop';
import {filter, map, take, tap} from 'rxjs';

export const authGuard: CanMatchFn = () => {
    const dispatcher = inject(Dispatcher);
    const sessionStore = inject(SessionStore);
    const router = inject(Router);

    if (sessionStore.status() === 'idle') {
        dispatcher.dispatch(userEvents.getUser());
    }

    const currentStatus = sessionStore.status();
    if (currentStatus === 'authenticated') return true;
    if (currentStatus === 'unauthenticated') {
        router.navigate(['/auth/login']);
        return false;
    }

    return toObservable(sessionStore.status).pipe(
        filter(status => status !== 'loading' && status !== 'idle'),
        take(1),
        tap(status => {
            if (status === 'unauthenticated') router.navigate(['/auth/login'])
        }),
        map(status => status === 'authenticated')
    );
};
