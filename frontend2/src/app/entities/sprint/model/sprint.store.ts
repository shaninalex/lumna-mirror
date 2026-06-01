import { createEntityAdapter, EntityState } from "@ngrx/entity";

import { SprintModel } from "./sprint.model";
import { createReducer, on } from "@ngrx/store";
import { actionSprintCreateSuccess, actionSprintSetList } from "./sprint.actions";

export interface SprintState extends EntityState<SprintModel> {}
export const sprintAdapter = createEntityAdapter<SprintModel>();
const initialState = sprintAdapter.getInitialState();

export const sprintReducer = createReducer(
    initialState,
    on(actionSprintSetList, (state, { sprints }) =>
        sprintAdapter.addMany(sprints, state)
    ),
    on(actionSprintCreateSuccess, (state, { sprint }) =>
        sprintAdapter.addOne(sprint, state)
    )
);
