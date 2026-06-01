import { createAction, props } from "@ngrx/store";
import {
    SprintCreateModel,
    SprintListQueryModel,
    SprintModel
} from "./sprint.model";
import { Error } from "@shared/models";

export const actionSprintGetList = createAction(
    "[Sprint] get list",
    props<{ query: SprintListQueryModel }>()
);

export const actionSprintSetList = createAction(
    "[Sprint] set list",
    props<{ sprints: SprintModel[] }>()
);

export const actionSprintSetListFailed = createAction(
    "[Sprint] set list failed",
    props<{ errors: Error[] }>()
);

export const actionSprintCreate = createAction(
    "[Sprint] create",
    props<{ data: SprintCreateModel }>()
);

export const actionSprintCreateSuccess = createAction(
    "[Sprint] create success",
    props<{ sprint: SprintModel }>()
);
export const actionSprintCreateFailed = createAction(
    "[Sprint] create failed",
    props<{ errors: Error[] }>()
);
