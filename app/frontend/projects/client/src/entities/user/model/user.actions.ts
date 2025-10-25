import { createAction, props } from "@ngrx/store"
import { Settings, UserModel } from "@client/entities/user"

export const UserGetAction = createAction("[user] get")

export const UserSetAction = createAction("[user] set", props<{ payload: UserModel | null }>())

export const UserUpdateSettingsAction = createAction("[user] update settings", props<{ payload: Settings }>())
