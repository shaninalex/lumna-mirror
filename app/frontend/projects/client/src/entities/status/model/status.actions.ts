import { createAction, props } from "@ngrx/store"
import { Status, StatusInput } from "./status.model"

export const StatusListGetActions = createAction("[status] get status list", props<{ projectId: number }>())

export const StatusListSetActions = createAction("[status] set status list", props<{ payload: Status[] }>())

export const StatusSetAction = createAction("[status] set status", props<{ payload: Status }>())

export const StatusCreateAction = createAction("[status] create status", props<{ payload: StatusInput; projectId: number }>())

export const StatusPatchAction = createAction(
	"[status] patch status",
	props<{ payload: StatusInput; projectId: number; statusId: number }>()
)

export const StatusPatchSuccessAction = createAction("[status] patch status success", props<{ payload: Status }>())

export const StatusDeleteAction = createAction("[status] delete status", props<{ projectId: number; statusId: number }>())

export const StatusDeleteSuccessAction = createAction("[status] delete status success", props<{ projectId: number; statusId: number }>())

export const StatusPatchSortAction = createAction(
	"[status] update status sort",
	props<{ projectId: number; payload: Record<number, number> }>()
)

export const StatusPatchSortSuccessAction = createAction("[status] update status sort done", props<{ payload: Status[] }>())
