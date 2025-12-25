import {catchError, map, Observable, tap, throwError} from 'rxjs';
import {DestroyRef, inject, Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {environment as env} from '@environments/environment.development';
import {Dispatcher, Events} from '@ngrx/signals/events';
import {sessionEvents} from '@core/store/session.store';
import {takeUntilDestroyed} from '@angular/core/rxjs-interop';


const REFRESH_INTERVAL_MS = 14.5 * 60 * 1000; // 14.5 min

@Injectable({
    providedIn: 'root'
})
export class TokenRefreshService {
    private readonly events = inject(Events);
    readonly dispatcher = inject(Dispatcher)

    constructor() {
        this.events
            .on(sessionEvents.refreshToken)
            .pipe(
                takeUntilDestroyed(),
                tap(() => {
                    console.log("restart timer")
                })
            )
            .subscribe()

        this.events
            .on(sessionEvents.logout, sessionEvents.sessionFailed)
            .pipe(
                takeUntilDestroyed(),
                tap(() => {
                    console.log("stop timer")
                })
            )
            .subscribe()
    }

    // initialize token refresh before it's expired.
    private triggerRefreshAccessToken(): void {
        this.dispatcher.dispatch(sessionEvents.refreshToken())
    }
}
