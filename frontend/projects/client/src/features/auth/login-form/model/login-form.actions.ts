import {createAction, props} from '@ngrx/store';
import {TFlowWithUI} from '@ui';

// LoginFormLoadedAction - login form loaded action
export const LoginFormLoadedAction = createAction(
    "[login form] loaded",
    props<{ form: TFlowWithUI }>(),
);

// LoginFormExpired - this action should trigger reloading
export const LoginFormExpired = createAction(
    "[login form] expired"
)
