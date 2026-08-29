import { createActionGroup, emptyProps, props } from '@ngrx/store';
import type { UserModel } from './user.model';

export const actionUser = createActionGroup({
    source: 'User',
    events: {
        /**
         * Dispatched to request loading the current user.
         *
         * Typically handled by an effect that performs an API call
         * to fetch the authenticated user's data.
         */
        get: emptyProps(),

        /**
         * Dispatched to clear all user-related state.
         *
         * Commonly used on logout, session expiration,
         * or when resetting the application state.
         */
        clear: emptyProps(),

        /**
         * Dispatched when user data has been successfully loaded.
         *
         * Updates the store with the provided user model,
         * usually as a result of a successful API response.
         *
         * @payload
         * - user: UserModel containing user details
         */
        set: props<{ user: UserModel }>(),
    },
});
