import {createAction} from '@ngrx/store';
import {UserModel, Settings} from '@client/entities/user';

export const GetUserAction = createAction(
    "[user] get",
)

export const SetUserAction = createAction(
    "[user] set",
    prompt<{ payload: UserModel }>(),
)

export const UpdateUserSettings = createAction(
    "[user] update settings",
    prompt<{ payload: Settings }>(),
)
