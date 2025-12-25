import {Subscription, tap, timer} from 'rxjs';
import {inject, Injectable} from '@angular/core';
import {Dispatcher, Events} from '@ngrx/signals/events';
import {sessionEvents} from '@core/store/session.store';
import {takeUntilDestroyed} from '@angular/core/rxjs-interop';
import {LS_REFRESH_STARTED_AT} from '@shared/global.const';

const REFRESH_INTERVAL_MS = 14.5 * 60 * 1000; // 14.5 min

@Injectable({
    providedIn: 'root'
})
export class TokenRefreshService {
    private readonly events = inject(Events);
    private readonly dispatcher = inject(Dispatcher);

    private refreshTimer$?: Subscription;

    constructor() {
        // Start timer when authenticated (login, bootstrap, or token refresh success)
        this.events
            .on(sessionEvents.authenticated)
            .pipe(takeUntilDestroyed())
            .subscribe(() => this.startRefreshTimer());

        // Stop timer on logout or session failure
        this.events
            .on(sessionEvents.logout, sessionEvents.sessionFailed)
            .pipe(takeUntilDestroyed())
            .subscribe(() => this.stopRefreshTimer());

        // Try to restore timer from previous session
        this.restoreTimerIfNeeded();
    }

    private startRefreshTimer(delayMs?: number, persistTime = true): void {
        // Stop existing timer but don't clear persisted time yet
        if (this.refreshTimer$) {
            this.refreshTimer$.unsubscribe();
            this.refreshTimer$ = undefined;
        }

        const delay = delayMs ?? REFRESH_INTERVAL_MS;
        if (persistTime) {
            this.persistStartTime();
        }

        console.log(`TokenRefresh: starting timer for ${Math.round(delay / 1000)}s`);

        this.refreshTimer$ = timer(delay).subscribe(() => {
            console.log('TokenRefresh: triggering refresh');
            this.dispatcher.dispatch(sessionEvents.refreshToken());
        });
    }

    private stopRefreshTimer(): void {
        if (this.refreshTimer$) {
            console.log('TokenRefresh: stopping timer');
            this.refreshTimer$.unsubscribe();
            this.refreshTimer$ = undefined;
        }
        this.clearPersistedTime();
    }

    private persistStartTime(): void {
        localStorage.setItem(LS_REFRESH_STARTED_AT, Date.now().toString());
    }

    private clearPersistedTime(): void {
        localStorage.removeItem(LS_REFRESH_STARTED_AT);
    }

    private restoreTimerIfNeeded(): void {
        const startedAtRaw = localStorage.getItem(LS_REFRESH_STARTED_AT);
        if (!startedAtRaw) return;

        const startedAt = Number(startedAtRaw);
        const elapsed = Date.now() - startedAt;
        const remaining = REFRESH_INTERVAL_MS - elapsed;

        if (remaining <= 0) {
            // Token might be expired or about to expire - refresh immediately after bootstrap
            console.log('TokenRefresh: timer expired during page load, will refresh after bootstrap');
            this.clearPersistedTime();
        } else {
            // Resume timer with remaining time (don't overwrite persisted time)
            console.log(`TokenRefresh: restoring timer with ${Math.round(remaining / 1000)}s remaining`);
            this.startRefreshTimer(remaining, false);
        }
    }
}
