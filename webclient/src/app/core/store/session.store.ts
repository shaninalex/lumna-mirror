import {patchState, signalStore, withState} from '@ngrx/signals';
import { withEventHandlers, Events } from '@ngrx/signals/events';
import { inject } from '@angular/core';
import { userEvents } from '@entities/user';
import { mapResponse } from '@ngrx/operators';
import { UserApi } from '@entities/user/api/user.service';
import { tap, switchMap } from 'rxjs';

type SessionState = {
    status: 'idle' | 'loading' | 'authenticated' | 'unauthenticated';
};

const initialState: SessionState = {
    status: 'idle',
};

export const SessionStore = signalStore(
    { providedIn: 'root' },
    withState(initialState),
    withEventHandlers((
        store,
        events = inject(Events),
        userApi = inject(UserApi),
    ) => ({
        bootstrap$: events
            .on(userEvents.getUser)
            .pipe(
                tap(() => patchState(store, { status: 'loading' })),
                switchMap(() =>
                    userApi.GetUser().pipe(
                        mapResponse({
                            next: (user) => userEvents.setUser(user),
                            error: () => userEvents.sessionFailed(),
                        })
                    )
                )
            ),

        authenticated$: events
            .on(userEvents.setUser)
            .pipe(
                tap(() => patchState(store, { status: 'authenticated' }))
            ),

        failed$: events
            .on(userEvents.sessionFailed)
            .pipe(
                tap(() => patchState(store, { status: 'unauthenticated' }))
            ),
    }))
);
