import { createAction, props } from "@ngrx/store";
import { ProjectModel } from "./project.model";

export const actionProjectGetList = createAction("[Project] get list");

export const actionProjectSetList = createAction(
    "[Project] set list",
    props<{ list: ProjectModel[] }>()
);
