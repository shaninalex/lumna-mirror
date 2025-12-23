import {signalStore, withState, patchState} from '@ngrx/signals';
import {inject} from '@angular/core';
import {userEvents} from '@entities/user';
import {UserApi} from '@entities/user/api/user.service';
import {switchMap, tap} from 'rxjs/operators';
import {Events, withEventHandlers} from '@ngrx/signals/events';
import {mapResponse} from '@ngrx/operators';

export type SessionStatus = 'idle' | 'loading' | 'authenticated' | 'unauthenticated';

type SessionState = {
    status: SessionStatus;
};

const initialState: SessionState = {
    status: 'idle',
};

export const SessionStore = signalStore(
    { providedIn: 'root' },
    withState<SessionState>(initialState),
    withEventHandlers((store, events = inject(Events), userApi = inject(UserApi)) => ({
        bootstrap$: events
            .on(userEvents.getUser)
            .pipe(
                tap(() => patchState(store, { status: 'loading' })),
                switchMap(() =>
                    userApi.GetUser().pipe(
                        mapResponse({
                            next: user => userEvents.setUser(user),
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

        unauthenticated$: events
            .on(userEvents.sessionFailed)
            .pipe(
                tap(() => patchState(store, { status: 'unauthenticated' }))
            ),
    }))
);
