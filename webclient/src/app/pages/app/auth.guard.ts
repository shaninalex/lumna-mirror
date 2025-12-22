import {CanActivateFn} from '@angular/router';
import {inject} from '@angular/core';
import {userEvents, UserStore} from '@entities/user';
import {Dispatcher} from '@ngrx/signals/events';

export const authGuard: CanActivateFn = (route, state) => {
    const userStore = inject(UserStore);
    const user = userStore.user;
    if (!user()) {
        const dispatcher = inject(Dispatcher)
        dispatcher.dispatch(userEvents.getUser());
    }
    return true;
};
