import { InvitationModel } from "@entities/invitation";
import { createAction, props } from "@ngrx/store";
import { Error } from "@shared/models";
import { OnboardingContinue, UserOnboardingModel } from "./onboarding.models";

/**
 * Dispatched when user opens invite user link with token
 * parameter.
 *
 * @payload
 * - token: invite token string
 */
export const actionOnboardingCreateInvite = createAction(
    "[Onboarding] create invite",
    props<{ payload: UserOnboardingModel }>()
);

/**
 * Dispatched when onboarding invitation attempt was success
 */
export const actionOnboardingCreateInviteSuccess = createAction(
    "[Onboarding] create invite success",
    props<{ invitation: InvitationModel }>()
);

/**
 * Dispatched when onboarding invitation attempt was failed
 */
export const actionOnboardingCreateInviteFailed = createAction(
    "[Onboarding] create invite failed",
    props<{ error: Error[] }>()
);

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
    "[Onboarding] Validate invite success",
    props<{ invitation: OnboardingContinue }>()
);

/**
 * Dispatched when invite link validation failed
 */
export const actionOnboardingValidateInviteFailed = createAction(
    "[Onboarding] Validate invite failed",
    props<{ error: Error }>()
);
