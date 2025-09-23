import {createAction, props} from '@ngrx/store';
import {UserModel, Settings} from '@client/entities/user';

export const GetUserAction = createAction(
    "[user] get",
)

export const SetUserAction = createAction(
    "[user] set",
    props<{ payload: UserModel | null }>(),
)

export const UpdateUserSettingsAction = createAction(
    "[user] update settings",
    props<{ payload: Settings }>(),
)
