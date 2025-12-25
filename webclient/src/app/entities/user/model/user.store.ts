import {patchState, signalStore, withHooks, withState} from '@ngrx/signals';
import {UserModel} from './user.model';
import {Events, withEventHandlers} from '@ngrx/signals/events';
import {inject} from '@angular/core';
import {userEvents} from '@entities/user';
import {tap} from 'rxjs';

type UserState = {
    user: UserModel | undefined;
};

const initialState: UserState = {
    user: undefined,
};

export const UserStore = signalStore(
    {providedIn: 'root'},
    withHooks({
        onInit() {
            console.log('UserStore initialized');
        },
    }),
    withState(initialState),
    withEventHandlers(
        (
            store,
            events = inject(Events),
        ) => ({
            setUser$: events
                .on(userEvents.setUser)
                .pipe(
                    tap(() => console.log("UserStore: set user")),
                    tap(eventData => patchState(store, {
                        user: eventData.payload,
                    }))
                ),

            clear$: events
                .on(userEvents.clear)
                .pipe(
                    tap(() => console.log("UserStore: clear")),
                    tap(() => patchState(store, { user: undefined }))
                )

        })
    )
);
