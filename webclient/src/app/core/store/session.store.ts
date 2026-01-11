import {patchState, signalStore, type, withState} from '@ngrx/signals';
import {inject} from '@angular/core';
import {userEvents, UserModel} from '@entities/user';
import {UserApi} from '@entities/user/api/user.service';
import {switchMap, tap} from 'rxjs/operators';
import {eventGroup, Events, withEventHandlers} from '@ngrx/signals/events';
import {mapResponse} from '@ngrx/operators';
import {LoginCredentials} from '@features/auth/login/model/login.model';
import {LoginService} from '@core/api/login.service';
import {map} from 'rxjs';
import {Router} from '@angular/router';

export const sessionEvents = eventGroup({
    source: 'Session',
    events: {
        authenticate: type<LoginCredentials>(),
        authenticated: type<UserModel>(),
        refreshToken: type<void>(),
        refreshSucceeded: type<void>(),
        sessionFailed: type<any>(), // TODO: replace any with auth error type
        logout: type<void>(),
    },
});

export type SessionStatus = 'idle' | 'loading' | 'authenticated' | 'unauthenticated';

type SessionState = {
    status: SessionStatus;
};

const initialState: SessionState = {status: 'idle'};

export const SessionStore = signalStore(
        {providedIn: 'root'},
        withState(initialState),
        withEventHandlers((
            store,
            events = inject(Events),
            userApi = inject(UserApi),
            loginApi = inject(LoginService),
            router = inject(Router),
        ) => ({
            authStart$: events
                .on(sessionEvents.authenticate)
                .pipe(
                    tap(() => patchState(store, {status: 'loading'})),
                    switchMap(event =>
                        loginApi.Login(event.payload).pipe(
                            mapResponse({
                                next: user => sessionEvents.authenticated(user),
                                error: event => sessionEvents.sessionFailed(event),
                            })
                        )
                    )
                ),

            refreshToken$: events
                .on(sessionEvents.refreshToken)
                .pipe(
                    tap(() => patchState(store, {status: 'loading'})),
                    switchMap(() =>
                        loginApi.Refresh().pipe(
                            mapResponse({
                                next: () => sessionEvents.refreshSucceeded(),
                                error: event => sessionEvents.sessionFailed(event),
                            })
                        )
                    )
                ),

            logout$: events
                .on(sessionEvents.logout)
                .pipe(
                    tap(() => patchState(store, {status: 'unauthenticated'})),
                    tap(() => console.log('SessionStore: process logout')),
                    switchMap(() =>
                        loginApi.Logout().pipe(
                            mapResponse({
                                next: () => router.navigate(['/auth/login']),
                                error: event => sessionEvents.sessionFailed(event),
                            }))
                    ),
                ),

            refreshSucceeded$:
                events
                    .on(sessionEvents.refreshSucceeded)
                    .pipe(
                        tap(() => patchState(store, {status: 'authenticated'})),
                    ),

            bootstrap$:
                events
                    .on(userEvents.getUser)
                    .pipe(
                        tap(() => patchState(store, {status: 'loading'})),
                        switchMap(() =>
                            userApi.GetUser().pipe(
                                tap(() => console.log('SessionStore: load user')),
                                mapResponse({
                                    next: user => sessionEvents.authenticated(user),
                                    error: event => sessionEvents.sessionFailed(event),
                                })
                            )
                        )
                    ),

            authenticated$:
                events.on(sessionEvents.authenticated).pipe(
                    tap(() => console.log('SessionStore: authenticated')),
                    tap(() => patchState(store, {status: 'authenticated'})),
                    map(e => userEvents.setUser(e.payload)),
                ),

            unauthenticated$:
                events.on(sessionEvents.sessionFailed).pipe(
                    tap(event => console.log("Notify about session failed error", event.payload)),
                    tap(() => router.navigate(['/auth/login'])),
                    tap(() => patchState(store, {status: 'unauthenticated'})),
                ),
        }))
    )
;
