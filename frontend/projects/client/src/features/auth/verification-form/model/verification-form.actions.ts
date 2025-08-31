import {createAction, props} from '@ngrx/store';
import {TFlowWithUI} from '@ui';

// VerificationFormLoadedAction - login form loaded action
export const VerificationFormLoadedAction = createAction(
    "[verification form] loaded",
    props<{ form: TFlowWithUI }>(),
);

// VerificationFormExpired - this action should trigger reloading
export const VerificationFormExpired = createAction(
    "[verification form] expired"
)
