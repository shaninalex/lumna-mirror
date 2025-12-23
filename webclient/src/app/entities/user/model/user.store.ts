import {patchState, signalStore, withState} from '@ngrx/signals';
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
    withState(initialState),
    withEventHandlers(
        (
            store,
            events = inject(Events),
        ) => ({
            setUser$: events
                .on(userEvents.setUser)
                .pipe(
                    tap(eventData => patchState(store, {
                        user: eventData.payload,
                    }))
                )

        })
    )
);
