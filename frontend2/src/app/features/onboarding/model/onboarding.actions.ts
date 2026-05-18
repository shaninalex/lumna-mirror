import { createAction, props } from "@ngrx/store";
import { Error } from "@shared/models";

/**
 * Dispatched when user opens invite user link with token
 * parameter.
 *
 * @payload
 * - token: invite token string
 */
export const actionOnboardingValidateInviteToken = createAction(
    "[Onboarding] Validate invite token",
    props<{ token: string }>()
);

/**
 * Dispatched when invite link validation successfull.
 * And user now proceed with invitation process ( provide user details )
 */
export const actionOnboardingValidateInviteSuccess = createAction(
    "[Onboarding] Validate invite success"
);

/**
 * Dispatched when invite link validation failed
 */
export const actionOnboardingValidateInviteFailed = createAction(
    "[Onboarding] Validate invite failed",
    props<{ error: Error }>()
);
