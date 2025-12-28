import {CanMatchFn} from '@angular/router';
import {inject} from '@angular/core';
import {userEvents} from '@entities/user';
import {Dispatcher} from '@ngrx/signals/events';
import {SessionStore} from '@core/store/session.store';
import {toObservable} from '@angular/core/rxjs-interop';
import {filter, map, take, tap} from 'rxjs';

export const authGuard: CanMatchFn = () => {
    const dispatcher = inject(Dispatcher);
    const sessionStore = inject(SessionStore);

    const status = sessionStore.status();

    // console.log('AuthGuard: ', status)
    // Already resolved - return immediately
    if (status === 'authenticated') return true;
    if (status === 'unauthenticated') return false;

    // Trigger bootstrap if idle
    if (status === 'idle') {
        // console.log('AuthGuard: ', 'dispatch get user')
        dispatcher.dispatch(userEvents.getUser());
    }

    // Wait for status to resolve
    return toObservable(sessionStore.status).pipe(
        // tap(() => console.log('AuthGuard: ', status)),
        filter(s => s === 'authenticated' || s === 'unauthenticated'),
        take(1),
        map(s => s === 'authenticated')
    );
};
