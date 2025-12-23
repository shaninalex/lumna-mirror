import { type } from '@ngrx/signals';
import { eventGroup } from '@ngrx/signals/events';

export const appEvents = eventGroup({
    source: 'App',
    events: {
        applicationReady: type<boolean>(),
    },
});
