import {patchState, signalStore, withState} from '@ngrx/signals';
import {UserModel} from './user.model';
import {Events, withEventHandlers} from '@ngrx/signals/events';
import {inject} from '@angular/core';
import {userEvents} from '@entities/user';
import {switchMap, tap} from 'rxjs';
import {mapResponse} from '@ngrx/operators';
import {UserApi} from '@entities/user/api/user.service';

type UserState = {
    user: UserModel | undefined;
    isLoading: boolean;
};

const initialState: UserState = {
    user: undefined,
    isLoading: false,
};

export const UserStore = signalStore(
    {providedIn: 'root'},
    withState(initialState),
    withEventHandlers(
        (
            store,
            events = inject(Events),
            userService = inject(UserApi)
        ) => ({
            userAuthenticated$: events
                .on(userEvents.authenticated)
                .pipe(
                    switchMap(() =>
                        userService.GetUser().pipe(
                            mapResponse({
                                next: (user) => userEvents.setUser(user),
                                error: error => console.log(error)
                            })
                        )
                    )
                ),
            setUser$: events
                .on(userEvents.setUser)
                .pipe(
                    tap(eventData => patchState(store, {user: eventData.payload, isLoading: false}))
                )

        })
    )
);
