import { inject } from '@angular/core';
import { patchState, signalStore, type, withState } from '@ngrx/signals';
import { eventGroup, Events, withEventHandlers } from '@ngrx/signals/events';
import { map, tap } from 'rxjs';

export const coreEvents = eventGroup({
    source: 'Core',
    events: {
        dataRequestFailed: type<string>(),
    },
});

type CoreState = {
    errors: string[];
};

const initialState: CoreState = {
    errors: [],
};

export const CoreStore = signalStore(
    { providedIn: 'root' },
    withState(initialState),
    withEventHandlers((store, events = inject(Events)) => ({
        error: events.on(coreEvents.dataRequestFailed).pipe(
            map((event) => event.payload),
            tap((message) => {
                patchState(store, (state) => ({
                    errors: [...state.errors, message],
                }));
            }),
        ),
    })),
);
