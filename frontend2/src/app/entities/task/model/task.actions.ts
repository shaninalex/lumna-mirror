import { createAction, props } from "@ngrx/store";
import { TaskListQueryModel, TaskModel } from "./task.model";
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
