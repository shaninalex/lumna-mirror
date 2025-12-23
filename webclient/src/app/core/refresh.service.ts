import {catchError, map, Observable, Subscription, switchMap, throwError, timer} from 'rxjs';
import {DestroyRef, inject, Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {environment as env} from '@environments/environment.development';

import {Events} from '@ngrx/signals/events';
import {userEvents} from '@entities/user';
import {takeUntilDestroyed} from '@angular/core/rxjs-interop';
import {LS_REFRESH_STARTED_AT} from '@shared/global.const';

const REFRESH_INTERVAL_MS = 14.5 * 60 * 1000; // 14.5 min

@Injectable({
    providedIn: 'root'
})
export class TokenRefreshService {
    private readonly http = inject(HttpClient);
    private readonly events = inject(Events);
    private readonly destroyRef = inject(DestroyRef);

    private refreshTimer$?: Subscription;


    constructor() {
        this.events
            .on(userEvents.setUser)
            .subscribe(() => {
                localStorage.setItem(
                    LS_REFRESH_STARTED_AT,
                    Date.now().toString()
                );
            });

        this.listenAuthEvents();
        this.restoreTimerIfNeeded();
    }

    private listenAuthEvents() {
        this.events
            .on(userEvents.setUser)
            .pipe(takeUntilDestroyed(this.destroyRef))
            .subscribe(() => {
                this.persistStartTime();
                this.startRefreshTimer();
            });

        // TODO: logout → clear timer + storage
    }

    private startRefreshTimer(delayMs?: number) {
        this.stopRefreshTimer();

        const delay = delayMs ?? REFRESH_INTERVAL_MS;

        this.refreshTimer$ = timer(delay)
            .pipe(switchMap(() => this.Refresh()))
            .subscribe({
                next: () => {
                    this.persistStartTime();
                    this.startRefreshTimer();
                },
                error: () => {
                    this.clearPersistedTime();
                    this.stopRefreshTimer();
                },
            });
    }

    private stopRefreshTimer() {
        this.refreshTimer$?.unsubscribe();
        this.refreshTimer$ = undefined;
    }

    private persistStartTime() {
        localStorage.setItem(
            LS_REFRESH_STARTED_AT,
            Date.now().toString()
        );
    }

    private clearPersistedTime() {
        localStorage.removeItem(LS_REFRESH_STARTED_AT);
    }

    private restoreTimerIfNeeded() {
        const startedAtRaw = localStorage.getItem(LS_REFRESH_STARTED_AT);
        if (!startedAtRaw) {
            return;
        }
        const startedAt = Number(startedAtRaw);
        const elapsed = Date.now() - startedAt;
        const remaining = REFRESH_INTERVAL_MS - elapsed;

        if (remaining <= 0) {
            this.persistStartTime();
            this.startRefreshTimer(0);
        } else {
            this.startRefreshTimer(remaining);
        }
    }

    public Refresh(): Observable<void> {
        return this.http
            .get(`${env.API_ROOT}/api/v1/auth/refresh`, {
                withCredentials: true,
            })
            .pipe(
                map(() => void 0),
                catchError((err) => throwError(() => err))
            );
    }
}
