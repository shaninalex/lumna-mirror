import {patchState, signalStore, withState} from '@ngrx/signals';
import {Events, withEventHandlers} from '@ngrx/signals/events';
import {inject} from '@angular/core';
import {tap} from 'rxjs';
import {appEvents} from '@core/store/app.events';

type ApplicationState = {
    ready: boolean
};

const initialState: ApplicationState = {
    ready: false
};

export const ApplicationStore = signalStore(
    { providedIn: 'root' },
    withState(initialState),
    withEventHandlers(
        (
            store,
            events = inject(Events),
        ) => ({
            ready$: events
                .on(appEvents.applicationReady)
                .pipe(
                    tap(eventData => patchState(store, {
                        ready: eventData.payload,
                    }))
                )
        })
    )
)

