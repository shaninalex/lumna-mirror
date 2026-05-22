import { createAction, props } from "@ngrx/store";
import { UserModel, AcceptInviteUserRegistrationForm } from "./user.model";

/**
 * Dispatched to request loading the current user.
 *
 * Typically handled by an effect that performs an API call
 * to fetch the authenticated user's data.
 */
export const actionUserGet = createAction("[User] Get");

/**
 * Dispatched to clear all user-related state.
 *
 * Commonly used on logout, session expiration,
 * or when resetting the application state.
 */
export const actionUserClear = createAction("[User] Clear");

/**
 * Dispatched when user data has been successfully loaded.
 *
 * Updates the store with the provided user model,
 * usually as a result of a successful API response.
 *
 * @payload
 * - user: UserModel containing user details
 */
export const actionUserSet = createAction(
    "[User] Set",
    props<{ user: UserModel }>()
);

/**
 * Dispatched when user submit invitation registration form.
 *
 * @payload
 * - user: AcceptInviteUserRegistrationForm invitation registration process form data
 */
export const actionUserInviteRegister = createAction(
    "[User] invite register",
    props<{ data: AcceptInviteUserRegistrationForm }>()
);
