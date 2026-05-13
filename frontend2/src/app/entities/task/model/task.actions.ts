import { createAction, props } from "@ngrx/store";
import { TaskModel } from "./task.model";

export const actionTaskGetList = createAction("[Task] get list");

export const actionTaskSetList = createAction(
    "[Task] set list",
    props<{ list: TaskModel[] }>()
);
