import { createEntityAdapter, EntityAdapter, EntityState } from "@ngrx/entity"
import { createReducer, on } from "@ngrx/store"
import {
	Status,
	StatusDeleteSuccessAction,
	StatusListSetActions,
	StatusPatchSortSuccessAction,
	StatusPatchSuccessAction,
	StatusSetAction,
} from "@client/entities/status"

export interface StatusState extends EntityState<Status> {}

export const statusAdapter: EntityAdapter<Status> = createEntityAdapter<Status>()
export const statusReducer = createReducer(
	statusAdapter.getInitialState(),
	on(StatusListSetActions, (state, action) => statusAdapter.addMany(action.payload, state)),
	on(StatusSetAction, (state, action) => statusAdapter.addOne(action.payload, state)),
	on(StatusPatchSuccessAction, (state, action) =>
		statusAdapter.updateOne(
			{
				id: action.payload.id,
				changes: action.payload,
			},
			state
		)
	),
	on(StatusDeleteSuccessAction, (state, action) => statusAdapter.removeOne(action.statusId, state)),
	on(StatusPatchSortSuccessAction, (state, action) =>
		statusAdapter.updateMany(
			action.payload.map(status => ({
				id: status.id,
				changes: status,
			})),
			state
		)
	)
)
