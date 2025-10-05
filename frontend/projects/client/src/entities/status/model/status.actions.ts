import {createAction, props} from '@ngrx/store';
import {Status, StatusInput} from './status.model';

export const GetStatusListActions = createAction(
    "[status] get status list",
    props<{ projectId: number }>(),
)

export const SetStatusListActions = createAction(
    "[status] set status list",
    props<{ payload: Status[] }>(),
)

export const SetStatusAction = createAction(
    "[status] set status",
    props<{ payload: Status }>(),
)

export const CreateStatusAction = createAction(
    "[status] create status",
    props<{ payload: StatusInput, projectId: number }>(),
)

export const PatchStatusAction = createAction(
    "[status] patch status",
    props<{ payload: StatusInput, projectId: number, statusId: number }>(),
)

export const PatchStatusSuccessAction = createAction(
    "[status] patch status success",
    props<{ payload: Status }>(),
)

export const DeleteStatusAction = createAction(
    "[status] delete status",
    props<{ projectId: number, statusId: number }>(),
)

export const DeleteStatusSuccessAction = createAction(
    "[status] delete status success",
    props<{ projectId: number, statusId: number }>(),
)

export const PatchStatusSortAction = createAction(
    "[status] update status sort",
    props<{ projectId: number, payload: Record<number, number>}>()
)

export const PatchStatusSortSuccessAction = createAction(
    "[status] update status sort done",
    props<{ payload: Status[] }>(),
)

