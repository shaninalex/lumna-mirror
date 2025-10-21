import {createAction, props} from '@ngrx/store';
import {AppError} from '@client/shared/common';

export const AppErrorAction = createAction(
    "[app] error",
    props<{ err: AppError }>(),
);
