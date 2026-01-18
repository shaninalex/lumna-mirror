import { patchState, signalStore, type, withMethods, withState } from '@ngrx/signals';
import { eventGroup } from '@ngrx/signals/events';

type SessionState = {
    isAuthenticated: boolean;
};

const initialState: SessionState = {
    isAuthenticated: false,
};

export const sessionEvents = eventGroup({
    source: 'Session',
    events: {
        authenticated: type<void>(),
        logout: type<void>(),
    },
});

export const SessionStore = signalStore(
    { providedIn: 'root' },
    withState(initialState),
    withMethods((store) => ({
        setAuthenticated(value: boolean): void {
            patchState(store, { isAuthenticated: value });
        },
    })),
);
