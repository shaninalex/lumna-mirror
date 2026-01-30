import { createAction, props } from '@ngrx/store';
import { UserModel } from '@entities/user';

export const actionSessionAuthenticate = createAction(
    '[Session] authenticate',
    props<{ user: UserModel }>(),
);
