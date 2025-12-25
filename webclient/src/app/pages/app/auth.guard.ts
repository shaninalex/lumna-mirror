import {CanMatchFn} from '@angular/router';
import {inject} from '@angular/core';
import {userEvents} from '@entities/user';
import {Dispatcher} from '@ngrx/signals/events';
import {SessionStore} from '@core/store/session.store';

export const authGuard: CanMatchFn = () => {
    const dispatcher = inject(Dispatcher);
    const sessionStore = inject(SessionStore);

    if (sessionStore.status() === 'idle') {
        dispatcher.dispatch(userEvents.getUser());
    }

    return sessionStore.status() === 'authenticated'
};
