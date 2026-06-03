import { createAction, props } from "@ngrx/store";
import { TaskCreateModel, TaskListQueryModel, TaskModel } from "./task.model";
import { Error } from "@shared/models";

export const actionTaskGetList = createAction(
    "[Task] get list",
    props<{ query: TaskListQueryModel }>()
);

export const actionTaskSetList = createAction(
    "[Task] set list",
    props<{ tasks: TaskModel[] }>()
);

export const actionTaskSetListFailed = createAction(
    "[Task] set list failed",
    props<{ errors: Error[] }>()
);

export const actionTaskCreate = createAction(
    "[Task] create",
    props<{ data: TaskCreateModel }>()
);

export const actionTaskCreateFailed = createAction(
    "[Task] create failed",
    props<{ errors: Error[] }>()
);

export const actionTaskCreateSuccess = createAction(
    "[Task] create success",
    props<{ task: TaskModel }>()
);
