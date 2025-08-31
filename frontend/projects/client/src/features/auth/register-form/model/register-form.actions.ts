import {createAction, props} from '@ngrx/store';
import {TFlowWithUI} from '@ui';

// RegisterFormLoadedAction - login form loaded action
export const RegisterFormLoadedAction = createAction(
    "[register form] loaded",
    props<{ form: TFlowWithUI }>(),
);

// RegisterFormExpired - this action should trigger reloading
export const RegisterFormExpired = createAction(
    "[register form] expired"
)
