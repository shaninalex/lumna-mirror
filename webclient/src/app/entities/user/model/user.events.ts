import { type } from '@ngrx/signals';
import {UserModel} from './user.model';
import { eventGroup } from '@ngrx/signals/events';

export const userEvents = eventGroup({
    source: 'User',
    events: {
        getUser: type<void>(),
        setUser: type<UserModel>(),
        sessionFailed: type<void>(),
    },
});
