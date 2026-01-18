import { patchState, signalStore, withMethods, withState } from '@ngrx/signals';

type SessionState = {
    isAuthenticated: boolean;
};

const initialState: SessionState = {
    isAuthenticated: false,
};

export const SessionStore = signalStore(
    { providedIn: 'root' },
    withState(initialState),
    withMethods((store) => ({
        setAuthenticated(value: boolean): void {
            patchState(store, { isAuthenticated: value });
        },
    })),
);
